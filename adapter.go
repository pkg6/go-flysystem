package flysystem

import (
	"fmt"
	"os"
	"strings"
)

type IAdapter interface {
	IFS
	// DiskName Default Disk Name
	DiskName() string
	// Clone Initialization parameters
	Clone() IAdapter
}

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
		return fmt.Sprintf("%s", strings.TrimPrefix(path, string(os.PathSeparator)))
	}
	return fmt.Sprintf("%s%s", a.prefix, strings.TrimPrefix(path, string(os.PathSeparator)))
}
