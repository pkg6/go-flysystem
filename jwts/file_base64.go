package jwts

import (
	"encoding/base64"
	"io"
	"os"
)

func FileBase64(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	fd, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(fd), nil
}
