package flysystem

import (
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	PathTypeFile      = "file"
	PathTypeDirectory = "directory"
	ModePublicString  = "public"
	ModePrivateString = "private"
	ModeFilePublic    = 0644
	ModeFilePrivate   = 0600
	ModeDirPublic     = 0755
	ModeDirPrivate    = 0700
)

var (
	FileModes = map[string]map[string]os.FileMode{
		PathTypeFile: {
			ModePublicString:  ModeFilePublic,
			ModePrivateString: ModeFilePrivate,
		},
		PathTypeDirectory: {
			ModePublicString:  ModeDirPublic,
			ModePrivateString: ModeDirPrivate,
		},
	}
)

type IFS interface {
	// Exists Determine if the file exists
	Exists(path string) (bool, error)
	// Size Get File Size
	Size(path string) (int64, error)
	// WriteReader write file content and return full path
	WriteReader(path string, reader io.Reader) (string, error)
	// Write  file content and return full path
	Write(path string, contents []byte) (string, error)
	// WriteStream Resource file write returns full path
	WriteStream(path, resource string) (string, error)
	// Update Update the file content and return the updated full path
	Update(path string, contents []byte) (string, error)
	// UpdateStream Return the updated full path based on resource file updates
	UpdateStream(path, resource string) (string, error)
	// Read Read file
	Read(path string) ([]byte, error)
	// Delete  Deleting files returns the number of deleted files
	Delete(path string) (int64, error)
	// DeleteDirectory Number of files deleted from the deleted directory
	DeleteDirectory(dirname string) (int64, error)
	// CreateDirectory create directory
	CreateDirectory(dirname string) error
	// MimeType Get File MimeType
	MimeType(path string) (string, error)
	// Move move file
	Move(source, destination string) (bool, error)
	// Copy copy file
	Copy(source, destination string) (bool, error)
}

type IFlysystem interface {
	IFS
	Extend(name string, adapter IAdapter) IFlysystem
	Disk(disk string) IFlysystem
	FindAdapter() IAdapter
}

type IAdapter interface {
	IFS
	// DiskName Default Disk Name
	DiskName() string
	// Clone Initialization parameters
	Clone() IAdapter
}

type Path struct {
	Path string
	Size int64
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
