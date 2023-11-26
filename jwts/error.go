package jwts

import "fmt"

type FileUploadError struct {
	Code int
	Msg  string
}

func NewError(code int, msg string) *FileUploadError {
	return &FileUploadError{Code: code, Msg: msg}
}

func (f *FileUploadError) Error() string {
	return fmt.Sprintf("code: %d, msg: %s", f.Code, f.Msg)
}
