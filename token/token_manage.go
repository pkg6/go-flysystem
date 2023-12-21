package token

import (
	"github.com/pkg6/go-flysystem"
	"io"
	"mime/multipart"
)

type Manage struct {
	token *JWTToken
	fs    *flysystem.Flysystem
}

func NewManage(token *JWTToken, flysystem *flysystem.Flysystem) *Manage {
	return &Manage{token: token, fs: flysystem}
}

func (t *Manage) BuildToken(aud, disk, bucket string) (*TokenResponse, error) {
	return t.token.BuildToken(aud, disk, bucket)
}

func (t *Manage) ParseToken(token string) (*FlysystemClaims, error) {
	return t.token.ParseToken(token)
}

func (t *Manage) WithTokenUploadMultipart(token, fileName string, file *multipart.FileHeader) (*Response, error) {
	return t.token.WithTokenUploadMultipart(t.fs, token, fileName, file)
}

func (t *Manage) WithTokenUploadReader(token, fileName string, reader io.Reader) (*Response, error) {
	return t.token.WithTokenUploadReader(t.fs, token, fileName, reader)
}

func (t *Manage) WithTokenUploadFilePath(token, fileName, filePath string) (*Response, error) {
	return t.token.WithTokenUploadFilePath(t.fs, token, fileName, filePath)
}

func (t *Manage) WithTokenUploadBase64(token, fileName, base64 string) (*Response, error) {
	return t.token.WithTokenUploadBase64(t.fs, token, fileName, base64)
}

func (t *Manage) WithTokenDelete(token, fileName string) (*Response, error) {
	return t.token.WithTokenDelete(t.fs, token, fileName)
}

func (t *Manage) Delete(disk, bucket, fileName string) (*Response, error) {
	return t.token.Delete(t.fs, disk, bucket, fileName)
}

func (t *Manage) UploadBase64(disk, bucket, fileName, base64Str string) (*Response, error) {
	return t.token.UploadBase64(t.fs, disk, bucket, fileName, base64Str)
}

func (t *Manage) UploadReader(disk, bucket, fileName string, reader io.Reader) (*Response, error) {
	return t.token.UploadReader(t.fs, disk, bucket, fileName, reader)
}
func (t *Manage) UploadByte(disk, bucket, fileName string, contents []byte) (*Response, error) {
	return t.token.UploadByte(t.fs, disk, bucket, fileName, contents)
}
