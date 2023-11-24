package flysystem

import (
	"github.com/zzqqw/gfs"
	"io"
	"net/url"
)

const (
	DiskNameLocal                   = gfs.DiskNameLocal
	DiskNameOSS                     = gfs.DiskNameOSS
	DiskNameCOS                     = gfs.DiskNameCOS
	DiskNameBOS                     = gfs.DiskNameBOS
	DiskNameGoogleCloudCloudStorage = gfs.DiskNameGoogleCloudCloudStorage
	DiskNameQiNiuKoDo               = gfs.DiskNameQiNiuKoDo
)

var (
	FileModes = gfs.FileModes
)

type IBFS interface {
	// Exists Determine if the file exists
	Exists(path string) (bool, error)
	// WriteReader write file content and return full path
	WriteReader(path string, reader io.Reader) (string, error)
	// Write  file content and return full path
	Write(path string, contents []byte) (string, error)
	// WriteStream Resource file write returns full path
	WriteStream(path, resource string) (string, error)
	// Read Read file
	Read(path string) ([]byte, error)
	// Delete  Deleting files returns the number of deleted files
	Delete(path string) (int64, error)
}

type IFS interface {
	IBFS
	// Size Get File Size
	Size(path string) (int64, error)
	// Update  the file content and return the updated full path
	Update(path string, contents []byte) (string, error)
	// UpdateStream Return the updated full path based on resource file updates
	UpdateStream(path, resource string) (string, error)
	// MimeType Get File MimeType
	MimeType(path string) (string, error)
	// Move move file
	Move(source, destination string) (bool, error)
	// Copy copy file
	Copy(source, destination string) (bool, error)
	URL(path string) (*url.URL, error)
}
type IAdapter interface {
	IFS
	// DiskName Default Disk Name
	DiskName() string

	GFSAdapter() gfs.IAdapter
}

type IConfig interface {
	New() IAdapter
}
