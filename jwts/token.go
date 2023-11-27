package jwts

import (
	"encoding/base64"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/pkg6/go-flysystem"
	"io"
	"mime/multipart"
	"net/http"
	"time"
)

type Token struct {
	Key        string
	Iss        string
	ExpTimeAdd time.Duration
}

func (t *Token) BuildToken(aud, disk, bucket string) (*TokenResponse, error) {
	resp := new(TokenResponse)
	nowT := time.Now().Unix()
	resp.ExpTime = time.Now().Add(t.ExpTimeAdd).Unix()
	resp.ExpireIn = resp.ExpTime - nowT
	claims := FlysystemClaims{Iss: t.Iss, Iat: nowT, Exp: resp.ExpTime, Aud: aud, Disk: disk, Bucket: bucket}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(t.Key))
	if err != nil {
		return nil, err
	}
	resp.Token = token
	return resp, nil
}

func (t *Token) ParseToken(token string) (*FlysystemClaims, error) {
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

func (t *Token) withToken(token string) (*FlysystemClaims, error) {
	if token == "" {
		return nil, NewError(http.StatusNonAuthoritativeInfo, "Token is empty")
	}
	customClaims, err := t.ParseToken(token)
	if err != nil {
		return nil, NewError(http.StatusForbidden, fmt.Sprintf("Token parsing failed err=%v", err))
	}
	return customClaims, nil
}

func (t *Token) WithTokenUploadMultipart(fs *flysystem.Flysystem, token, fileName string, file *multipart.FileHeader) (*UploadResponse, error) {
	fileOpen, err := file.Open()
	if err != nil {
		return nil, NewError(http.StatusLengthRequired, fmt.Sprintf("file.Open() err=%v", err))
	}
	defer fileOpen.Close()
	return t.WithTokenUploadReader(fs, token, fileName, fileOpen)
}

func (t *Token) WithTokenUploadReader(fs *flysystem.Flysystem, token, fileName string, reader io.Reader) (*UploadResponse, error) {
	customClaims, err := t.withToken(token)
	if err != nil {
		return nil, err
	}
	return t.UploadReader(fs, customClaims.Disk, customClaims.Bucket, fileName, reader)
}

func (t *Token) WithTokenUploadFilePath(fs *flysystem.Flysystem, token, fileName, filePath string) (*UploadResponse, error) {
	fileBase64, err := FileBase64(filePath)
	if err != nil {
		return nil, NewError(http.StatusPreconditionFailed, fmt.Sprintf("Base64 parsing failed err=%v", err))
	}
	return t.WithTokenUploadBase64(fs, token, fileName, fileBase64)
}

func (t *Token) WithTokenUploadBase64(fs *flysystem.Flysystem, token, fileName, base64 string) (*UploadResponse, error) {
	customClaims, err := t.withToken(token)
	if err != nil {
		return nil, err
	}
	return t.UploadBase64(fs, customClaims.Disk, customClaims.Bucket, fileName, base64)
}

func (t *Token) UploadBase64(fs *flysystem.Flysystem, disk, bucket, fileName, base64Str string) (*UploadResponse, error) {
	fileBase64, _ := base64.StdEncoding.DecodeString(base64Str)
	return t.UploadByte(fs, disk, bucket, fileName, fileBase64)
}

func (t *Token) UploadReader(fs *flysystem.Flysystem, disk, bucket, fileName string, reader io.Reader) (*UploadResponse, error) {
	resp := &UploadResponse{Object: fileName, Bucket: bucket, Disk: disk}
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
func (t *Token) UploadByte(fs *flysystem.Flysystem, disk, bucket, fileName string, contents []byte) (*UploadResponse, error) {
	resp := &UploadResponse{Object: fileName, Bucket: bucket, Disk: disk}
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
