package fscos

import (
	"github.com/zzqqw/gfs/cosfs"
	"io"
	"net/url"
	"sync"

	"github.com/pkg6/go-flysystem"
	"github.com/zzqqw/gfs"
)

type FSCos struct {
	gfs.AbstractAdapter
	Config *Config
	lock   *sync.Mutex
}

func New(config *Config) flysystem.IAdapter {
	return NewCOS(config)
}
func NewCOS(config *Config) *FSCos {
	f := &FSCos{Config: config, lock: &sync.Mutex{}}
	f.SetPathPrefix(f.Config.PathPrefix)
	return f
}

func (a *FSCos) GFSAdapter() gfs.IAdapter {
	return cosfs.NewCOS(&cosfs.Config{
		CDN:       a.Config.CDN,
		BucketURL: a.Config.BucketURL,
		SecretID:  a.Config.SecretID,
		SecretKey: a.Config.SecretKey,
	})
}
func (a *FSCos) URL(path string) (*url.URL, error) {
	path = a.ApplyPathPrefix(path)
	return a.GFSAdapter().URL(path)
}
func (a *FSCos) Exists(path string) (bool, error) {
	path = a.ApplyPathPrefix(path)
	return a.GFSAdapter().Exist(path)
}

func (a *FSCos) WriteReader(path string, reader io.Reader) (string, error) {
	path = a.ApplyPathPrefix(path)
	err := a.GFSAdapter().WriteReader(path, reader)
	return path, err
}

func (a *FSCos) Write(path string, contents []byte) (string, error) {

	path = a.ApplyPathPrefix(path)
	err := a.GFSAdapter().Write(path, contents)
	return path, err
}

func (a *FSCos) WriteStream(path, resource string) (string, error) {

	path = a.ApplyPathPrefix(path)
	err := a.GFSAdapter().WriteStream(path, resource)
	return path, err
}

func (a *FSCos) Read(path string) ([]byte, error) {

	path = a.ApplyPathPrefix(path)
	return a.GFSAdapter().Read(path)
}

func (a *FSCos) Delete(path string) (int64, error) {

	path = a.ApplyPathPrefix(path)
	return a.GFSAdapter().Delete(path)
}

func (a *FSCos) Size(path string) (int64, error) {

	path = a.ApplyPathPrefix(path)
	return a.GFSAdapter().Size(path)
}

func (a *FSCos) Update(path string, contents []byte) (string, error) {

	path = a.ApplyPathPrefix(path)
	err := a.GFSAdapter().Update(path, contents)
	return path, err
}

func (a *FSCos) UpdateStream(path, resource string) (string, error) {

	path = a.ApplyPathPrefix(path)
	err := a.GFSAdapter().UpdateStream(path, resource)
	return path, err
}

func (a *FSCos) MimeType(path string) (string, error) {

	path = a.ApplyPathPrefix(path)
	return a.GFSAdapter().MimeType(path)
}

func (a *FSCos) Move(source, destination string) (bool, error) {

	source = a.ApplyPathPrefix(source)
	destination = a.ApplyPathPrefix(destination)
	return a.GFSAdapter().Move(source, destination)
}

func (a *FSCos) Copy(source, destination string) (bool, error) {

	source = a.ApplyPathPrefix(source)
	destination = a.ApplyPathPrefix(destination)
	return a.GFSAdapter().Copy(source, destination)
}

func (a *FSCos) DiskName() string {
	return flysystem.DiskNameCOS
}
