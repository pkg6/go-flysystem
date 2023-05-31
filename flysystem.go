package flysystem

import (
	"fmt"
)

type Flysystem struct {
	disk         string
	diskAdapters map[string]IAdapter
	diskNames    []string
}

func New() IFlysystem {
	return &Flysystem{diskAdapters: make(map[string]IAdapter)}
}

func NewAdapters(adapters ...IAdapter) IFlysystem {
	f := New()
	for _, adapter := range adapters {
		f.Extend(adapter.DiskName(), adapter)
	}
	return f
}

// Extend 扩展
func (f *Flysystem) Extend(name string, adapter IAdapter) IFlysystem {
	f.diskAdapters[name] = adapter
	f.diskNames = append(f.diskNames, name)
	return f
}

func (f *Flysystem) Disk(disk string) IFlysystem {
	return &Flysystem{
		disk:         disk,
		diskAdapters: f.diskAdapters,
		diskNames:    f.diskNames,
	}
}

// FindAdapter Find Adapter
func (f *Flysystem) FindAdapter() IAdapter {
	var disk string
	if f.disk != "" {
		disk = f.disk
	} else if len(f.diskNames) > 0 {
		disk = f.diskNames[0]
	}
	if adapter, ok := f.diskAdapters[disk]; ok {
		return adapter
	}
	panic(fmt.Sprintf("Unable to find %s disk", disk))
}
func (f *Flysystem) Exists(path string) (bool, error) {
	return f.FindAdapter().Exists(path)
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
func (f *Flysystem) CreateDirectory(dirname string) error {
	return f.FindAdapter().CreateDirectory(dirname)
}
func (f *Flysystem) DeleteDirectory(dirname string) (int64, error) {
	return f.FindAdapter().DeleteDirectory(dirname)
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
