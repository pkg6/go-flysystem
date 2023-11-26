package flysystem

import (
	"fmt"
	"io"
	"net/url"
	"reflect"
	"strings"
	"sync"

	"github.com/zzqqw/gfs"
)

type Flysystem struct {
	disk         string
	diskAdapters map[string]IAdapter
	diskNames    []string
	l            *sync.Mutex
}

func NewConfig(config any) (*Flysystem, error) {
	fs := &Flysystem{l: &sync.Mutex{}, diskAdapters: make(map[string]IAdapter)}
	err := fs.ExtendConfigPtr(config)
	return fs, err
}
func New() *Flysystem {
	return &Flysystem{diskAdapters: make(map[string]IAdapter), l: &sync.Mutex{}}
}

func NewAdapters(adapters ...IAdapter) *Flysystem {
	f := &Flysystem{diskAdapters: make(map[string]IAdapter), l: &sync.Mutex{}}
	for _, adapter := range adapters {
		f.Extend(adapter)
	}
	return f
}
func (f *Flysystem) ExtendConfigPtr(config any) error {
	v := reflect.ValueOf(config)
	t := reflect.TypeOf(config)
	if t.Kind() == reflect.Ptr {
		for i := 0; i < v.Elem().NumField(); i++ {
			e := v.Elem().Field(i)
			if !e.IsZero() {
				if fsConfig, ok := e.Interface().(IConfig); ok {
					var diskName string
					gfsName := t.Elem().Field(i).Tag.Get(gfs.ConfigPtrTag)
					if gfsName == "" {
						diskName = t.Elem().Field(i).Name
					} else {
						split := strings.Split(gfsName, ",")
						diskName = split[0]
						if len(split) > 2 && split[1] == gfs.ConfigPtrSplitTagDefaultDisk {
							if f.disk == "" {
								f.disk = diskName
							}
						}
					}
					f.Extend(fsConfig.New(), diskName)
				}
			}
		}
		return nil
	}
	return fmt.Errorf("the data type is incorrect %v", config)
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

// Disks 获取注册所有的驱动
func (f *Flysystem) Disks() []string {
	return f.diskNames
}

// DiskExist 判断驱动是否存在
func (f *Flysystem) DiskExist(disk string) bool {
	_, ok := f.diskAdapters[disk]
	return ok
}
func (f *Flysystem) Disk(disk string) string {
	if disk != "" {
		f.disk = disk
	}
	if f.disk == "" {
		f.disk = f.diskNames[0]
	}
	return f.disk
}

// Adapter Find Adapter
func (f *Flysystem) Adapter(disk string) (IAdapter, error) {
	if adapter, ok := f.diskAdapters[f.Disk(disk)]; ok {
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
