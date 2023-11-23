package flysystem

import (
	"fmt"
	"io"
	"net/url"
	"sync"

	"github.com/zzqqw/gfs"
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

type Flysystem struct {
	disk         string
	diskAdapters map[string]IAdapter
	diskNames    []string
	l            *sync.Mutex
}

func New() *Flysystem {
	return &Flysystem{diskAdapters: make(map[string]IAdapter), l: &sync.Mutex{}}
}

func NewAdapters(adapters ...IAdapter) *Flysystem {
	f := New()
	for _, adapter := range adapters {
		f.Extend(adapter)
	}
	return f
}

// Extend 扩展
func (f *Flysystem) Extend(adapter IAdapter, names ...string) *Flysystem {
	f.l.Lock()
	defer f.l.Unlock()
	name := adapter.DiskName()
	if len(names) > 0 {
		name = names[0]
	}
	f.diskAdapters[name] = adapter
	f.diskNames = append(f.diskNames, name)
	return f
}

// DiskGet 获取注册所有的驱动
func (f *Flysystem) DiskGet() []string {
	return f.diskNames
}

// DiskExist 判断驱动是否存在
func (f *Flysystem) DiskExist(disk string) bool {
	_, ok := f.diskAdapters[disk]
	return ok
}

// Adapter Find Adapter
func (f *Flysystem) Adapter(disk string) (IAdapter, error) {
	if disk != "" {
		f.disk = disk
	} else {
		f.disk = f.diskNames[0]
	}
	if adapter, ok := f.diskAdapters[f.disk]; ok {
		return adapter, nil
	}
	return nil, fmt.Errorf("unable to find【%s】disk", f.disk)
}

func (f *Flysystem) GFSAdapter(disk string) (gfs.IAdapter, error) {
	adapter, err := f.Adapter(disk)
	if err != nil {
		return nil, err
	}
	return adapter.GFSAdapter(), nil
}

func (f *Flysystem) URL(path string) (*url.URL, error) {
	adapter, err := f.Adapter("")
	if err != nil {
		return nil, err
	}
	return adapter.URL(path)
}

func (f *Flysystem) Exists(path string) (bool, error) {
	adapter, err := f.Adapter("")
	if err != nil {
		return false, err
	}
	return adapter.Exists(path)
}

func (f *Flysystem) WriteReader(path string, reader io.Reader) (string, error) {
	adapter, err := f.Adapter("")
	if err != nil {
		return "", err
	}
	return adapter.WriteReader(path, reader)
}

func (f *Flysystem) Write(path string, contents []byte) (string, error) {
	adapter, err := f.Adapter("")
	if err != nil {
		return "", err
	}
	return adapter.Write(path, contents)
}

func (f *Flysystem) WriteStream(path, resource string) (string, error) {
	adapter, err := f.Adapter("")
	if err != nil {
		return "", err
	}
	return adapter.WriteStream(path, resource)
}

func (f *Flysystem) Update(path string, contents []byte) (string, error) {
	adapter, err := f.Adapter("")
	if err != nil {
		return "", err
	}
	return adapter.Update(path, contents)
}

func (f *Flysystem) UpdateStream(path, resource string) (string, error) {
	adapter, err := f.Adapter("")
	if err != nil {
		return "", err
	}
	return adapter.UpdateStream(path, resource)
}

func (f *Flysystem) Read(path string) ([]byte, error) {
	adapter, err := f.Adapter("")
	if err != nil {
		return nil, err
	}
	return adapter.Read(path)
}

func (f *Flysystem) Delete(path string) (int64, error) {
	adapter, err := f.Adapter("")
	if err != nil {
		return 0, err
	}
	return adapter.Delete(path)
}

func (f *Flysystem) MimeType(path string) (string, error) {
	adapter, err := f.Adapter("")
	if err != nil {
		return "", err
	}
	return adapter.MimeType(path)
}

func (f *Flysystem) Size(path string) (int64, error) {
	adapter, err := f.Adapter("")
	if err != nil {
		return 0, err
	}
	return adapter.Size(path)
}

func (f *Flysystem) Move(source, destination string) (bool, error) {
	adapter, err := f.Adapter("")
	if err != nil {
		return false, err
	}
	return adapter.Move(source, destination)
}

func (f *Flysystem) Copy(source, destination string) (bool, error) {
	adapter, err := f.Adapter("")
	if err != nil {
		return false, err
	}
	return adapter.Copy(source, destination)
}
