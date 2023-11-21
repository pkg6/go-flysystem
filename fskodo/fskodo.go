package fskodo

import (
	"github.com/zzqqw/gfs/kodofs"
	"io"
	"net/url"
	"sync"

	"github.com/pkg6/go-flysystem"
	"github.com/qiniu/go-sdk/v7/storage"
	"github.com/zzqqw/gfs"
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

func (a *FSKodo) GFSAdapter() gfs.IAdapter {
	return kodofs.NewKoDo(&kodofs.Config{
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
	return a.GFSAdapter().URL(path)
}
func (a *FSKodo) Exists(path string) (bool, error) {
	a.init()
	path = a.ApplyPathPrefix(path)
	return a.GFSAdapter().Exist(path)
}

func (a *FSKodo) WriteReader(path string, reader io.Reader) (string, error) {
	a.init()
	path = a.ApplyPathPrefix(path)
	err := a.GFSAdapter().WriteReader(path, reader)
	return path, err
}

func (a *FSKodo) Write(path string, contents []byte) (string, error) {
	a.init()
	path = a.ApplyPathPrefix(path)
	err := a.GFSAdapter().Write(path, contents)
	return path, err
}

func (a *FSKodo) WriteStream(path, resource string) (string, error) {
	a.init()
	path = a.ApplyPathPrefix(path)
	err := a.GFSAdapter().WriteStream(path, resource)
	return path, err
}

func (a *FSKodo) Read(path string) ([]byte, error) {
	a.init()
	path = a.ApplyPathPrefix(path)
	return a.GFSAdapter().Read(path)
}

func (a *FSKodo) Delete(path string) (int64, error) {
	a.init()
	path = a.ApplyPathPrefix(path)
	return a.GFSAdapter().Delete(path)
}

func (a *FSKodo) Size(path string) (int64, error) {
	a.init()
	path = a.ApplyPathPrefix(path)
	return a.GFSAdapter().Size(path)
}

func (a *FSKodo) Update(path string, contents []byte) (string, error) {
	a.init()
	path = a.ApplyPathPrefix(path)
	err := a.GFSAdapter().Update(path, contents)
	return path, err
}

func (a *FSKodo) UpdateStream(path, resource string) (string, error) {
	a.init()
	path = a.ApplyPathPrefix(path)
	err := a.GFSAdapter().UpdateStream(path, resource)
	return path, err
}

func (a *FSKodo) MimeType(path string) (string, error) {
	a.init()
	path = a.ApplyPathPrefix(path)
	return a.GFSAdapter().MimeType(path)
}

func (a *FSKodo) Move(source, destination string) (bool, error) {
	a.init()
	source = a.ApplyPathPrefix(source)
	destination = a.ApplyPathPrefix(destination)
	return a.GFSAdapter().Move(source, destination)
}

func (a *FSKodo) Copy(source, destination string) (bool, error) {
	a.init()
	source = a.ApplyPathPrefix(source)
	destination = a.ApplyPathPrefix(destination)
	return a.GFSAdapter().Copy(source, destination)
}

func (a *FSKodo) DiskName() string {
	return flysystem.DiskNameQiNiuKoDo
}
