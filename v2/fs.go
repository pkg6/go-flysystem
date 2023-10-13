package v2

import (
	"fmt"
	"io"
	"sync"
)

type FsManage struct {
	disk         string
	disks        []string
	diskAdapters map[string]IAdapter
	l            *sync.Mutex
}

func NewConfig(config IConfig) IFSManage {
	fs := &FsManage{disk: config.Disk(), l: &sync.Mutex{}}
	adapters := config.Adapters()
	for s, adapter := range adapters {
		fs.Extend(adapter, s)
	}
	return fs
}

func New() IFSManage {
	fs := &FsManage{
		diskAdapters: make(map[string]IAdapter),
		l:            &sync.Mutex{},
	}
	return fs
}

// Extend 扩展
func (f *FsManage) Extend(adapter IAdapter, names ...string) IFSManage {
	f.l.Lock()
	defer f.l.Unlock()
	var name string
	if len(names) > 0 {
		name = names[0]
	} else {
		name = adapter.DiskName()
	}
	f.disks = append(f.disks, name)
	f.diskAdapters[name] = adapter
	return f
}

func (f *FsManage) Disk(disk string) IFSManage {
	return &FsManage{
		disk:         disk,
		diskAdapters: f.diskAdapters,
	}
}

// FindAdapter Find Adapter
func (f *FsManage) FindAdapter() IAdapter {
	disk := f.disk
	if disk == "" {
		f.disk = f.disks[0]
	}
	if adapter, ok := f.diskAdapters[f.disk]; ok {
		return adapter
	}
	panic(fmt.Sprintf("Unable to find %s disk", f.disk))
}

func (f *FsManage) Exist(path string) (bool, error) {
	return f.FindAdapter().Exist(path)
}

func (f *FsManage) WriteReader(path string, reader io.Reader) error {
	return f.FindAdapter().WriteReader(path, reader)
}

func (f *FsManage) Write(path string, contents []byte) error {
	return f.FindAdapter().Write(path, contents)
}

func (f *FsManage) WriteStream(path, resource string) error {
	return f.FindAdapter().WriteStream(path, resource)
}

func (f *FsManage) Update(path string, contents []byte) error {
	return f.FindAdapter().Update(path, contents)
}

func (f *FsManage) UpdateStream(path, resource string) error {
	return f.FindAdapter().UpdateStream(path, resource)
}

func (f *FsManage) Read(path string) ([]byte, error) {
	return f.FindAdapter().Read(path)
}

func (f *FsManage) Delete(path string) (int64, error) {
	return f.FindAdapter().Delete(path)
}

func (f *FsManage) MimeType(path string) (string, error) {
	return f.FindAdapter().MimeType(path)
}

func (f *FsManage) Size(path string) (int64, error) {
	return f.FindAdapter().Size(path)
}

func (f *FsManage) Move(source, destination string) (bool, error) {
	return f.FindAdapter().Move(source, destination)
}

func (f *FsManage) Copy(source, destination string) (bool, error) {
	return f.FindAdapter().Copy(source, destination)
}
