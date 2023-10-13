package v2

import (
	"io"
	"os"
)

const (
	DiskNameLocal                   = "LOCAL"
	DiskNameOSS                     = "OSS"
	DiskNameCOS                     = "COS"
	DiskNameBOS                     = "BOS"
	DiskNameGoogleCloudCloudStorage = "CloudStorage"
	DiskNameQiNiuKoDo               = "KODO"

	HeaderGetLength      = "content-length"
	HeaderGetContentType = "content-type"

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
	// Exist Determine if the file exists
	Exist(path string) (bool, error)
	// WriteReader write file content and return full path
	WriteReader(path string, reader io.Reader) error
	// Write  file content and return full path
	Write(path string, contents []byte) error
	// WriteStream Resource file write returns full path
	WriteStream(path, resource string) error
	// Read Read file
	Read(path string) ([]byte, error)
	// Delete  Deleting files returns the number of deleted files
	Delete(path string) (int64, error)
	// Size Get File Size
	Size(path string) (int64, error)
	// Update  the file content and return the updated full path
	Update(path string, contents []byte) error
	// UpdateStream Return the updated full path based on resource file updates
	UpdateStream(path, resource string) error
	// MimeType Get File MimeType
	MimeType(path string) (string, error)
	// Move move file
	Move(source, destination string) (bool, error)
	// Copy copy file
	Copy(source, destination string) (bool, error)
}

type IAdapter interface {
	IFS
	// DiskName Default Disk Name
	DiskName() string
}

type IAdapterConfig interface {
	New() IAdapter
}

type IFSManage interface {
	IFS
	Extend(adapter IAdapter, names ...string) IFSManage
	Disk(disk string) IFSManage
	FindAdapter() IAdapter
}

type IConfig interface {
	Disk() string
	Adapters() map[string]IAdapter
}
