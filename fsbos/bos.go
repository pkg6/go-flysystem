package fsbos

import (
	"github.com/pkg6/go-flysystem"
	"github.com/zzqqw/gfs"
	"github.com/zzqqw/gfs/bosfs"
	"io"
	"net/url"
	"sync"
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
	gfs.AbstractAdapter
	Config *Config
	lock   *sync.Mutex
}

func New(config *Config) flysystem.IAdapter {
	a := &FSBos{Config: config, lock: &sync.Mutex{}}
	a.SetPathPrefix(a.Config.PathPrefix)
	return a
}

func (a *FSBos) GFSAdapter() gfs.IAdapter {
	return bosfs.NewBOS(&bosfs.Config{
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
	return a.GFSAdapter().URL(path)
}
func (a *FSBos) Exists(path string) (bool, error) {
	path = a.ApplyPathPrefix(path)
	return a.GFSAdapter().Exist(path)
}

func (a *FSBos) WriteReader(path string, reader io.Reader) (string, error) {
	path = a.ApplyPathPrefix(path)
	err := a.GFSAdapter().WriteReader(path, reader)
	return path, err
}

func (a *FSBos) Write(path string, contents []byte) (string, error) {
	path = a.ApplyPathPrefix(path)
	err := a.GFSAdapter().Write(path, contents)
	return path, err
}

func (a *FSBos) WriteStream(path, resource string) (string, error) {
	path = a.ApplyPathPrefix(path)
	err := a.GFSAdapter().WriteStream(path, resource)
	return path, err
}

func (a *FSBos) Read(path string) ([]byte, error) {
	path = a.ApplyPathPrefix(path)
	return a.GFSAdapter().Read(path)
}

func (a *FSBos) Delete(path string) (int64, error) {
	path = a.ApplyPathPrefix(path)
	return a.GFSAdapter().Delete(path)
}

func (a *FSBos) Size(path string) (int64, error) {
	path = a.ApplyPathPrefix(path)
	return a.GFSAdapter().Size(path)
}

func (a *FSBos) Update(path string, contents []byte) (string, error) {
	path = a.ApplyPathPrefix(path)
	err := a.GFSAdapter().Update(path, contents)
	return path, err
}

func (a *FSBos) UpdateStream(path, resource string) (string, error) {
	path = a.ApplyPathPrefix(path)
	err := a.GFSAdapter().UpdateStream(path, resource)
	return path, err
}

func (a *FSBos) MimeType(path string) (string, error) {
	path = a.ApplyPathPrefix(path)
	return a.GFSAdapter().MimeType(path)
}

func (a *FSBos) Move(source, destination string) (bool, error) {
	source = a.ApplyPathPrefix(source)
	destination = a.ApplyPathPrefix(destination)
	return a.GFSAdapter().Move(source, destination)
}

func (a *FSBos) Copy(source, destination string) (bool, error) {
	source = a.ApplyPathPrefix(source)
	destination = a.ApplyPathPrefix(destination)
	return a.GFSAdapter().Copy(source, destination)
}

func (a *FSBos) DiskName() string {
	return flysystem.DiskNameBOS
}
