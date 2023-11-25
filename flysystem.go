package flysystem

import (
	"fmt"
	"io"
	"net/url"
	"reflect"
	"sync"

	"github.com/zzqqw/gfs"
)

type Flysystem struct {
	disk         string
	diskAdapters map[string]IAdapter
	diskNames    []string
	l            *sync.Mutex
}

// FlysystemExtend 通过配置来加载节点
func FlysystemExtend(fs *Flysystem, c any) {
	v := reflect.ValueOf(c)
	t := reflect.TypeOf(c)
	var disks []string
	diskConfigs := map[string]IConfig{}
	for i := 0; i < v.Elem().NumField(); i++ {
		e := v.Elem().Field(i)
		if !e.IsZero() {
			if fsConfig, ok := e.Interface().(IConfig); ok {
				name := t.Elem().Field(i).Name
				disks = append(disks, name)
				diskConfigs[name] = fsConfig
			}
		}
	}
	for _, disk := range disks {
		if config, ok := diskConfigs[disk]; ok {
			fs.Extend(config.New(), disk)
		}
	}
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
