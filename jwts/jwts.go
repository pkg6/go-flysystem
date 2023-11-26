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

type Jwts struct {
	Key string
	Iss string
}

func (j *Jwts) BuildToken(exp time.Duration, aud, disk, bucket string) (*TokenResponse, error) {
	resp := new(TokenResponse)
	t := time.Now().Unix()
	resp.ExpTime = time.Now().Add(exp).Unix()
	resp.ExpireIn = resp.ExpTime - t
	claims := FlysystemClaims{Iss: j.Iss, Iat: t, Exp: resp.ExpTime, Aud: aud, Disk: disk, Bucket: bucket}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(j.Key))
	if err != nil {
		return nil, err
	}
	resp.Token = token
	return resp, nil
}

func (j *Jwts) ParseToken(token string) (*FlysystemClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &FlysystemClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.Key), nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*FlysystemClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

func (j *Jwts) withToken(token string) (*FlysystemClaims, error) {
	if token == "" {
		return nil, NewError(http.StatusNonAuthoritativeInfo, "Token is empty")
	}
	customClaims, err := j.ParseToken(token)
	if err != nil {
		return nil, NewError(http.StatusForbidden, fmt.Sprintf("Token parsing failed err=%v", err))
	}
	return customClaims, nil
}

func (j *Jwts) WithTokenUploadMultipart(fs *flysystem.Flysystem, token, fileName string, file *multipart.FileHeader) error {
	fileOpen, err := file.Open()
	if err != nil {
		return NewError(http.StatusLengthRequired, fmt.Sprintf("file.Open() err=%v", err))
	}
	defer fileOpen.Close()
	return j.WithTokenUploadReader(fs, token, fileName, fileOpen)
}

func (j *Jwts) WithTokenUploadReader(fs *flysystem.Flysystem, token, fileName string, reader io.Reader) error {
	customClaims, err := j.withToken(token)
	if err != nil {
		return err
	}
	return j.UploadReader(fs, customClaims.Disk, customClaims.Bucket, fileName, reader)
}

func (j *Jwts) WithTokenUploadFilePath(fs *flysystem.Flysystem, token, fileName, filePath string) error {
	fileBase64, err := FileBase64(filePath)
	if err != nil {
		return NewError(http.StatusPreconditionFailed, fmt.Sprintf("Base64 parsing failed err=%v", err))
	}
	return j.WithTokenUploadBase64(fs, token, fileName, fileBase64)
}

func (j *Jwts) WithTokenUploadBase64(fs *flysystem.Flysystem, token, fileName, base64 string) error {
	customClaims, err := j.withToken(token)
	if err != nil {
		return err
	}
	return j.UploadBase64(fs, customClaims.Disk, customClaims.Bucket, fileName, base64)
}

func (j *Jwts) UploadBase64(fs *flysystem.Flysystem, disk, bucket, fileName, base64Str string) error {
	fileBase64, _ := base64.StdEncoding.DecodeString(base64Str)
	return j.UploadByte(fs, disk, bucket, fileName, fileBase64)
}

func (j *Jwts) UploadReader(fs *flysystem.Flysystem, disk, bucket, fileName string, reader io.Reader) error {
	gfs, err := fs.GFSAdapter(disk)
	if err != nil {
		return NewError(http.StatusUseProxy, fmt.Sprintf("GFSAdapter Driver 【%s】 not found", disk))
	}
	gfs.Bucket(bucket)
	if err := gfs.WriteReader(fileName, reader); err != nil {
		return NewError(http.StatusNoContent, err.Error())
	}
	return nil
}
func (j *Jwts) UploadByte(fs *flysystem.Flysystem, disk, bucket, fileName string, contents []byte) error {
	gfs, err := fs.GFSAdapter(disk)
	if err != nil {
		return NewError(http.StatusUseProxy, fmt.Sprintf("GFSAdapter Driver 【%s】 not found", disk))
	}
	gfs.Bucket(bucket)
	if err := gfs.Write(fileName, contents); err != nil {
		return NewError(http.StatusNoContent, err.Error())
	}
	return nil
}
