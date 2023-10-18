package fsbos

import (
	"github.com/pkg6/go-flysystem"
	"github.com/pkg6/go-flysystem/v2"
	fsbos2 "github.com/pkg6/go-flysystem/v2/fsbos"
	"io"
	"net/url"
)

type Config struct {
	CDN              string
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

func (a *FSBos) Adapter() *fsbos2.Adapter {
	return fsbos2.NewBOS(&fsbos2.Config{
		CDN:              a.Config.CDN,
		Ak:               a.Config.Ak,
		Sk:               a.Config.Sk,
		Endpoint:         a.Config.Endpoint,
		RedirectDisabled: a.Config.RedirectDisabled,
		Bucket:           a.Config.Bucket,
	})
}
func (a *FSBos) URL(path string) (*url.URL, error) {
	path = a.ApplyPathPrefix(path)
	return a.Adapter().URL(path)
}
func (a *FSBos) Exists(path string) (bool, error) {
	path = a.ApplyPathPrefix(path)
	return a.Adapter().Exist(path)
}

func (a *FSBos) WriteReader(path string, reader io.Reader) (string, error) {
	path = a.ApplyPathPrefix(path)
	err := a.Adapter().WriteReader(path, reader)
	return path, err
}

func (a *FSBos) Write(path string, contents []byte) (string, error) {
	path = a.ApplyPathPrefix(path)
	err := a.Adapter().Write(path, contents)
	return path, err
}

func (a *FSBos) WriteStream(path, resource string) (string, error) {
	path = a.ApplyPathPrefix(path)
	err := a.Adapter().WriteStream(path, resource)
	return path, err
}

func (a *FSBos) Read(path string) ([]byte, error) {
	path = a.ApplyPathPrefix(path)
	return a.Adapter().Read(path)
}

func (a *FSBos) Delete(path string) (int64, error) {
	path = a.ApplyPathPrefix(path)
	return a.Adapter().Delete(path)
}

func (a *FSBos) Size(path string) (int64, error) {
	path = a.ApplyPathPrefix(path)
	return a.Adapter().Size(path)
}

func (a *FSBos) Update(path string, contents []byte) (string, error) {
	path = a.ApplyPathPrefix(path)
	err := a.Adapter().Update(path, contents)
	return path, err
}

func (a *FSBos) UpdateStream(path, resource string) (string, error) {
	path = a.ApplyPathPrefix(path)
	err := a.Adapter().UpdateStream(path, resource)
	return path, err
}

func (a *FSBos) MimeType(path string) (string, error) {
	path = a.ApplyPathPrefix(path)
	return a.Adapter().MimeType(path)
}

func (a *FSBos) Move(source, destination string) (bool, error) {
	source = a.ApplyPathPrefix(source)
	destination = a.ApplyPathPrefix(destination)
	return a.Adapter().Move(source, destination)
}

func (a *FSBos) Copy(source, destination string) (bool, error) {
	source = a.ApplyPathPrefix(source)
	destination = a.ApplyPathPrefix(destination)
	return a.Adapter().Copy(source, destination)
}

func (a *FSBos) DiskName() string {
	return flysystem.DiskNameBOS
}
