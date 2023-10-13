package fsbos

import (
	"github.com/pkg6/go-flysystem"
	"github.com/pkg6/go-flysystem/v2"
	fsbos2 "github.com/pkg6/go-flysystem/v2/fsbos"
	"io"
)

type Config struct {
	Ak               string
	Sk               string
	Endpoint         string
	RedirectDisabled bool
	Bucket           string
	PathPrefix       string
}
type FSBos struct {
	v2.AbstractAdapter
	Config *Config
}

func New(config *Config) flysystem.IAdapter {
	return &FSBos{Config: config}
}

func (f *FSBos) Adapter() *fsbos2.Adapter {
	return fsbos2.NewBOS(&fsbos2.Config{
		Ak:               f.Config.Ak,
		Sk:               f.Config.Sk,
		Endpoint:         f.Config.Endpoint,
		RedirectDisabled: f.Config.RedirectDisabled,
		Bucket:           f.Config.Bucket,
	})
}

func (f *FSBos) Exists(path string) (bool, error) {
	path = f.ApplyPathPrefix(path)
	return f.Adapter().Exist(path)
}

func (f *FSBos) WriteReader(path string, reader io.Reader) (string, error) {
	path = f.ApplyPathPrefix(path)
	err := f.Adapter().WriteReader(path, reader)
	return path, err
}

func (f *FSBos) Write(path string, contents []byte) (string, error) {
	path = f.ApplyPathPrefix(path)
	err := f.Adapter().Write(path, contents)
	return path, err
}

func (f *FSBos) WriteStream(path, resource string) (string, error) {
	path = f.ApplyPathPrefix(path)
	err := f.Adapter().WriteStream(path, resource)
	return path, err
}

func (f *FSBos) Read(path string) ([]byte, error) {
	path = f.ApplyPathPrefix(path)
	return f.Adapter().Read(path)
}

func (f *FSBos) Delete(path string) (int64, error) {
	path = f.ApplyPathPrefix(path)
	return f.Adapter().Delete(path)
}

func (f *FSBos) Size(path string) (int64, error) {
	path = f.ApplyPathPrefix(path)
	return f.Adapter().Size(path)
}

func (f *FSBos) Update(path string, contents []byte) (string, error) {
	path = f.ApplyPathPrefix(path)
	err := f.Adapter().Update(path, contents)
	return path, err
}

func (f *FSBos) UpdateStream(path, resource string) (string, error) {
	path = f.ApplyPathPrefix(path)
	err := f.Adapter().UpdateStream(path, resource)
	return path, err
}

func (f *FSBos) MimeType(path string) (string, error) {
	path = f.ApplyPathPrefix(path)
	return f.Adapter().MimeType(path)
}

func (f *FSBos) Move(source, destination string) (bool, error) {
	source = f.ApplyPathPrefix(source)
	destination = f.ApplyPathPrefix(destination)
	return f.Adapter().Move(source, destination)
}

func (f *FSBos) Copy(source, destination string) (bool, error) {
	source = f.ApplyPathPrefix(source)
	destination = f.ApplyPathPrefix(destination)
	return f.Adapter().Copy(source, destination)
}

func (f *FSBos) DiskName() string {
	return flysystem.DiskNameBOS
}
