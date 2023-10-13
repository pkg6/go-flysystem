package v2

import (
	"fmt"
	"os"
	"strings"
)

type AbstractAdapter struct {
	prefix string
}

func (a *AbstractAdapter) SetPathPrefix(prefix string) {
	if prefix != "" {
		a.prefix = fmt.Sprintf("%s%s", prefix, string(os.PathSeparator))
	}
}

func (a *AbstractAdapter) ApplyPathPrefix(path string) string {
	if a.prefix == "" {
		return strings.TrimPrefix(path, string(os.PathSeparator))
	}
	return fmt.Sprintf("%s%s", a.prefix, strings.TrimPrefix(path, string(os.PathSeparator)))
}
