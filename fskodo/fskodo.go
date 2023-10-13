package fskodo

import (
	"github.com/pkg6/go-flysystem"
	"github.com/pkg6/go-flysystem/v2"
	fskodo2 "github.com/pkg6/go-flysystem/v2/fskodo"
	"github.com/qiniu/go-sdk/v7/storage"
	"io"
)

type Config struct {
	AccessKey, SecretKey string
	Bucket               string
	Policy               *storage.PutPolicy
	Config               *storage.Config
	PathPrefix           string
}
type FSKodo struct {
	v2.AbstractAdapter
	Config *Config
}

func New(config *Config) flysystem.IAdapter {
	return &FSKodo{Config: config}
}

func (f *FSKodo) Adapter() *fskodo2.Adapter {
	return fskodo2.NewKoDo(&fskodo2.Config{
		AccessKey: f.Config.AccessKey,
		SecretKey: f.Config.SecretKey,
		Bucket:    f.Config.Bucket,
		Policy:    f.Config.Policy,
		Config:    f.Config.Config,
	})
}

func (f *FSKodo) Exists(path string) (bool, error) {
	path = f.ApplyPathPrefix(path)
	return f.Adapter().Exist(path)
}

func (f *FSKodo) WriteReader(path string, reader io.Reader) (string, error) {
	path = f.ApplyPathPrefix(path)
	err := f.Adapter().WriteReader(path, reader)
	return path, err
}

func (f *FSKodo) Write(path string, contents []byte) (string, error) {
	path = f.ApplyPathPrefix(path)
	err := f.Adapter().Write(path, contents)
	return path, err
}

func (f *FSKodo) WriteStream(path, resource string) (string, error) {
	path = f.ApplyPathPrefix(path)
	err := f.Adapter().WriteStream(path, resource)
	return path, err
}

func (f *FSKodo) Read(path string) ([]byte, error) {
	path = f.ApplyPathPrefix(path)
	return f.Adapter().Read(path)
}

func (f *FSKodo) Delete(path string) (int64, error) {
	path = f.ApplyPathPrefix(path)
	return f.Adapter().Delete(path)
}

func (f *FSKodo) Size(path string) (int64, error) {
	path = f.ApplyPathPrefix(path)
	return f.Adapter().Size(path)
}

func (f *FSKodo) Update(path string, contents []byte) (string, error) {
	path = f.ApplyPathPrefix(path)
	err := f.Adapter().Update(path, contents)
	return path, err
}

func (f *FSKodo) UpdateStream(path, resource string) (string, error) {
	path = f.ApplyPathPrefix(path)
	err := f.Adapter().UpdateStream(path, resource)
	return path, err
}

func (f *FSKodo) MimeType(path string) (string, error) {
	path = f.ApplyPathPrefix(path)
	return f.Adapter().MimeType(path)
}

func (f *FSKodo) Move(source, destination string) (bool, error) {
	source = f.ApplyPathPrefix(source)
	destination = f.ApplyPathPrefix(destination)
	return f.Adapter().Move(source, destination)
}

func (f *FSKodo) Copy(source, destination string) (bool, error) {
	source = f.ApplyPathPrefix(source)
	destination = f.ApplyPathPrefix(destination)
	return f.Adapter().Copy(source, destination)
}

func (f *FSKodo) DiskName() string {
	return flysystem.DiskNameQiNiuKoDo
}
