package flysystem

import (
	"github.com/pkg6/gfs"
	"io"
	"os"
)

func DecodeBase64(s string) ([]byte, error) {
	return gfs.Base64Decode(s)
}

func EncodeBase64(src []byte) string {
	return gfs.Base64Encode(src)
}

// OpenFileBase64 文件路径转base64
func OpenFileBase64(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	fd, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}
	return EncodeBase64(fd), nil
}

// SaveFileBase64 文件路径转base64
func SaveFileBase64(path string, data []byte) error {
	data, err := DecodeBase64(string(data))
	if err != nil {
		return err
	}
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	_, err = file.Write(data)
	if err != nil {
		return err
	}
	return nil
}
