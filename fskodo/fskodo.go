package fskodo

import (
	"github.com/pkg6/go-flysystem"
	"github.com/pkg6/go-flysystem/gfs"
	fskodo2 "github.com/pkg6/go-flysystem/gfs/fskodo"
	"github.com/qiniu/go-sdk/v7/storage"
	"io"
	"net/url"
	"sync"
)

type Config struct {
	CDN                  string
	AccessKey, SecretKey string
	Bucket               string
	Policy               *storage.PutPolicy
	Config               *storage.Config
	PathPrefix           string
}
type FSKodo struct {
	gfs.AbstractAdapter
	Config *Config
	lock   *sync.Mutex
}

func New(config *Config) flysystem.IAdapter {
	return &FSKodo{Config: config}
}
func (a *FSKodo) init() {
	if a.lock == nil {
		a.lock = &sync.Mutex{}
	}
	a.SetPathPrefix(a.Config.PathPrefix)
}

func (a *FSKodo) Adapter() *fskodo2.Adapter {
	return fskodo2.NewKoDo(&fskodo2.Config{
		CDN:       a.Config.CDN,
		AccessKey: a.Config.AccessKey,
		SecretKey: a.Config.SecretKey,
		Bucket:    a.Config.Bucket,
		Policy:    a.Config.Policy,
		Config:    a.Config.Config,
	})
}
func (a *FSKodo) URL(path string) (*url.URL, error) {
	a.init()
	path = a.ApplyPathPrefix(path)
	return a.Adapter().URL(path)
}
func (a *FSKodo) Exists(path string) (bool, error) {
	a.init()
	path = a.ApplyPathPrefix(path)
	return a.Adapter().Exist(path)
}

func (a *FSKodo) WriteReader(path string, reader io.Reader) (string, error) {
	a.init()
	path = a.ApplyPathPrefix(path)
	err := a.Adapter().WriteReader(path, reader)
	return path, err
}

func (a *FSKodo) Write(path string, contents []byte) (string, error) {
	a.init()
	path = a.ApplyPathPrefix(path)
	err := a.Adapter().Write(path, contents)
	return path, err
}

func (a *FSKodo) WriteStream(path, resource string) (string, error) {
	a.init()
	path = a.ApplyPathPrefix(path)
	err := a.Adapter().WriteStream(path, resource)
	return path, err
}

func (a *FSKodo) Read(path string) ([]byte, error) {
	a.init()
	path = a.ApplyPathPrefix(path)
	return a.Adapter().Read(path)
}

func (a *FSKodo) Delete(path string) (int64, error) {
	a.init()
	path = a.ApplyPathPrefix(path)
	return a.Adapter().Delete(path)
}

func (a *FSKodo) Size(path string) (int64, error) {
	a.init()
	path = a.ApplyPathPrefix(path)
	return a.Adapter().Size(path)
}

func (a *FSKodo) Update(path string, contents []byte) (string, error) {
	a.init()
	path = a.ApplyPathPrefix(path)
	err := a.Adapter().Update(path, contents)
	return path, err
}

func (a *FSKodo) UpdateStream(path, resource string) (string, error) {
	a.init()
	path = a.ApplyPathPrefix(path)
	err := a.Adapter().UpdateStream(path, resource)
	return path, err
}

func (a *FSKodo) MimeType(path string) (string, error) {
	a.init()
	path = a.ApplyPathPrefix(path)
	return a.Adapter().MimeType(path)
}

func (a *FSKodo) Move(source, destination string) (bool, error) {
	a.init()
	source = a.ApplyPathPrefix(source)
	destination = a.ApplyPathPrefix(destination)
	return a.Adapter().Move(source, destination)
}

func (a *FSKodo) Copy(source, destination string) (bool, error) {
	a.init()
	source = a.ApplyPathPrefix(source)
	destination = a.ApplyPathPrefix(destination)
	return a.Adapter().Copy(source, destination)
}

func (a *FSKodo) DiskName() string {
	return flysystem.DiskNameQiNiuKoDo
}
