package token

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/pkg6/go-flysystem"
	"io"
	"mime/multipart"
	"net/http"
	"time"
)

const (
	defaultSubject = "go-flysystem-token"
	defaultIssuer  = "go-flysystem"
)

type JWTToken struct {
	Key       string
	ExpiresIn time.Duration
}

func (t *JWTToken) BuildToken(aud, disk, bucket string) (*TokenResponse, error) {
	resp := new(TokenResponse)
	n := time.Now()
	claims := FlysystemClaims{Disk: disk, Bucket: bucket}
	claims.Subject = defaultSubject
	claims.Issuer = defaultIssuer
	claims.Audience = aud
	claims.Id = uuid.New().String()
	claims.IssuedAt = n.Unix()
	if t.ExpiresIn != time.Duration(0) {
		resp.ExpTime = n.Add(t.ExpiresIn).Unix()
		resp.ExpireIn = t.ExpiresIn.Seconds()
		claims.ExpiresAt = resp.ExpTime
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(t.Key))
	if err != nil {
		return nil, err
	}
	resp.Token = token
	return resp, nil
}

func (t *JWTToken) ParseToken(token string) (*FlysystemClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &FlysystemClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(t.Key), nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*FlysystemClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

func (t *JWTToken) withToken(token string) (*FlysystemClaims, error) {
	if token == "" {
		return nil, NewError(http.StatusNonAuthoritativeInfo, "Token is empty")
	}
	customClaims, err := t.ParseToken(token)
	if err != nil {
		return nil, NewError(http.StatusForbidden, fmt.Sprintf("Token parsing failed err=%v", err))
	}
	return customClaims, nil
}

func (t *JWTToken) WithTokenUploadMultipart(fs *flysystem.Flysystem, token, fileName string, file *multipart.FileHeader) (*Response, error) {
	fileOpen, err := file.Open()
	if err != nil {
		return nil, NewError(http.StatusLengthRequired, fmt.Sprintf("file.Open() err=%v", err))
	}
	defer fileOpen.Close()
	return t.WithTokenUploadReader(fs, token, fileName, fileOpen)
}

func (t *JWTToken) WithTokenUploadReader(fs *flysystem.Flysystem, token, fileName string, reader io.Reader) (*Response, error) {
	customClaims, err := t.withToken(token)
	if err != nil {
		return nil, err
	}
	return t.UploadReader(fs, customClaims.Disk, customClaims.Bucket, fileName, reader)
}

func (t *JWTToken) WithTokenUploadFilePath(fs *flysystem.Flysystem, token, fileName, filePath string) (*Response, error) {
	fileBase64, err := flysystem.OpenFileBase64(filePath)
	if err != nil {
		return nil, NewError(http.StatusPreconditionFailed, fmt.Sprintf("Base64 parsing failed err=%v", err))
	}
	return t.WithTokenUploadBase64(fs, token, fileName, fileBase64)
}

func (t *JWTToken) WithTokenUploadBase64(fs *flysystem.Flysystem, token, fileName, base64 string) (*Response, error) {
	customClaims, err := t.withToken(token)
	if err != nil {
		return nil, err
	}
	return t.UploadBase64(fs, customClaims.Disk, customClaims.Bucket, fileName, base64)
}

func (t *JWTToken) WithTokenDelete(fs *flysystem.Flysystem, token, fileName string) (*Response, error) {
	customClaims, err := t.withToken(token)
	if err != nil {
		return nil, err
	}
	return t.Delete(fs, customClaims.Disk, customClaims.Bucket, fileName)
}

func (t *JWTToken) Delete(fs *flysystem.Flysystem, disk, bucket, fileName string) (*Response, error) {
	resp := &Response{Object: fileName, Bucket: bucket, Disk: disk}
	gfs, err := fs.GFSAdapter(disk)
	if err != nil {
		return resp, NewError(http.StatusUseProxy, fmt.Sprintf("GFSAdapter Driver 【%s】 not found", disk))
	}
	gfs.Bucket(bucket)
	if _, err := gfs.Delete(fileName); err != nil {
		return resp, NewError(http.StatusNotFound, err.Error())
	}
	return resp, nil
}

func (t *JWTToken) UploadBase64(fs *flysystem.Flysystem, disk, bucket, fileName, base64Str string) (*Response, error) {
	fileBase64, _ := flysystem.DecodeBase64(base64Str)
	return t.UploadByte(fs, disk, bucket, fileName, fileBase64)
}

func (t *JWTToken) UploadReader(fs *flysystem.Flysystem, disk, bucket, fileName string, reader io.Reader) (*Response, error) {
	resp := &Response{Object: fileName, Bucket: bucket, Disk: disk}
	gfs, err := fs.GFSAdapter(disk)
	if err != nil {
		return resp, NewError(http.StatusUseProxy, fmt.Sprintf("GFSAdapter Driver 【%s】 not found", disk))
	}
	gfs.Bucket(bucket)
	if err := gfs.WriteReader(fileName, reader); err != nil {
		return resp, NewError(http.StatusNoContent, err.Error())
	}
	return resp, nil
}
func (t *JWTToken) UploadByte(fs *flysystem.Flysystem, disk, bucket, fileName string, contents []byte) (*Response, error) {
	resp := &Response{Object: fileName, Bucket: bucket, Disk: disk}
	gfs, err := fs.GFSAdapter(disk)
	if err != nil {
		return resp, NewError(http.StatusUseProxy, fmt.Sprintf("GFSAdapter Driver 【%s】 not found", disk))
	}
	gfs.Bucket(bucket)
	if err := gfs.Write(fileName, contents); err != nil {
		return resp, NewError(http.StatusNoContent, err.Error())
	}
	return resp, nil
}
