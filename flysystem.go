package flysystem

import (
	"fmt"
	"io"
	"os"
	"sync"
)

const (
	DiskNameOSS                     = "ALiYunOSS"
	DiskNameGoogleCloudCloudStorage = "GoogleCloudCloudStorage"
	DiskNameLocal                   = "Local"

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
}

type IFlysystem interface {
	IFS
	Extend(adapter IAdapter, names ...string) IFlysystem
	Disk(disk string) IFlysystem
	FindAdapter() IAdapter
}

type Flysystem struct {
	disk         string
	diskAdapters map[string]IAdapter
	l            *sync.Mutex
}

func New() IFlysystem {
	return &Flysystem{
		diskAdapters: make(map[string]IAdapter),
		l:            &sync.Mutex{},
	}
}

func NewAdapters(adapters ...IAdapter) IFlysystem {
	f := New()
	for _, adapter := range adapters {
		f.Extend(adapter)
	}
	return f
}

// Extend 扩展
func (f *Flysystem) Extend(adapter IAdapter, names ...string) IFlysystem {
	f.l.Lock()
	defer f.l.Unlock()
	name := adapter.DiskName()
	if len(names) > 0 {
		name = names[0]
	}
	f.diskAdapters[name] = adapter
	return f
}

func (f *Flysystem) Disk(disk string) IFlysystem {
	return &Flysystem{
		disk:         disk,
		diskAdapters: f.diskAdapters,
	}
}

// FindAdapter Find Adapter
func (f *Flysystem) FindAdapter() IAdapter {
	var disk string
	if f.disk != "" {
		disk = f.disk
	}
	if adapter, ok := f.diskAdapters[disk]; ok {
		return adapter
	}
	panic(fmt.Sprintf("Unable to find %s disk", disk))
}

func (f *Flysystem) Exists(path string) (bool, error) {
	return f.FindAdapter().Exists(path)
}

func (f *Flysystem) WriteReader(path string, reader io.Reader) (string, error) {
	return f.FindAdapter().WriteReader(path, reader)
}

func (f *Flysystem) Write(path string, contents []byte) (string, error) {
	return f.FindAdapter().Write(path, contents)
}

func (f *Flysystem) WriteStream(path, resource string) (string, error) {
	return f.FindAdapter().WriteStream(path, resource)
}

func (f *Flysystem) Update(path string, contents []byte) (string, error) {
	return f.FindAdapter().Update(path, contents)
}

func (f *Flysystem) UpdateStream(path, resource string) (string, error) {
	return f.FindAdapter().UpdateStream(path, resource)
}

func (f *Flysystem) Read(path string) ([]byte, error) {
	return f.FindAdapter().Read(path)
}

func (f *Flysystem) Delete(path string) (int64, error) {
	return f.FindAdapter().Delete(path)
}

func (f *Flysystem) MimeType(path string) (string, error) {
	return f.FindAdapter().MimeType(path)
}

func (f *Flysystem) Size(path string) (int64, error) {
	return f.FindAdapter().Size(path)
}

func (f *Flysystem) Move(source, destination string) (bool, error) {
	return f.FindAdapter().Move(source, destination)
}

func (f *Flysystem) Copy(source, destination string) (bool, error) {
	return f.FindAdapter().Copy(source, destination)
}
