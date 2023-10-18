package fscloudstorage

import (
	"github.com/pkg6/go-flysystem"
	"github.com/pkg6/go-flysystem/gfs"
	fscloudstorage2 "github.com/pkg6/go-flysystem/gfs/fscloudstorage"
	"google.golang.org/api/option"
	"io"
	"net/url"
	"sync"
	"time"
)

type Config struct {
	CDN             string
	Bucket          string
	WithTimeout     time.Duration
	CredentialsFile string
	Option          []option.ClientOption
	PathPrefix      string
}
type FSCloudStorage struct {
	gfs.AbstractAdapter
	Config *Config
	lock   *sync.Mutex
}

func New(config *Config) flysystem.IAdapter {
	return &FSCloudStorage{Config: config}
}
func (a *FSCloudStorage) init() {
	if a.lock == nil {
		a.lock = &sync.Mutex{}
	}
	a.SetPathPrefix(a.Config.PathPrefix)
}

func (a *FSCloudStorage) Adapter() *fscloudstorage2.Adapter {
	return fscloudstorage2.NewGCS(&fscloudstorage2.Config{
		CDN:             a.Config.CDN,
		Bucket:          a.Config.Bucket,
		WithTimeout:     a.Config.WithTimeout,
		CredentialsFile: a.Config.CredentialsFile,
		Option:          a.Config.Option,
	})
}
func (a *FSCloudStorage) URL(path string) (*url.URL, error) {
	a.init()
	path = a.ApplyPathPrefix(path)
	return a.Adapter().URL(path)
}
func (a *FSCloudStorage) Exists(path string) (bool, error) {
	a.init()
	path = a.ApplyPathPrefix(path)
	return a.Adapter().Exist(path)
}

func (a *FSCloudStorage) WriteReader(path string, reader io.Reader) (string, error) {
	a.init()
	path = a.ApplyPathPrefix(path)
	err := a.Adapter().WriteReader(path, reader)
	return path, err
}

func (a *FSCloudStorage) Write(path string, contents []byte) (string, error) {
	a.init()
	path = a.ApplyPathPrefix(path)
	err := a.Adapter().Write(path, contents)
	return path, err
}

func (a *FSCloudStorage) WriteStream(path, resource string) (string, error) {
	a.init()
	path = a.ApplyPathPrefix(path)
	err := a.Adapter().WriteStream(path, resource)
	return path, err
}

func (a *FSCloudStorage) Read(path string) ([]byte, error) {
	a.init()
	path = a.ApplyPathPrefix(path)
	return a.Adapter().Read(path)
}

func (a *FSCloudStorage) Delete(path string) (int64, error) {
	a.init()
	path = a.ApplyPathPrefix(path)
	return a.Adapter().Delete(path)
}

func (a *FSCloudStorage) Size(path string) (int64, error) {
	a.init()
	path = a.ApplyPathPrefix(path)
	return a.Adapter().Size(path)
}

func (a *FSCloudStorage) Update(path string, contents []byte) (string, error) {
	a.init()
	path = a.ApplyPathPrefix(path)
	err := a.Adapter().Update(path, contents)
	return path, err
}

func (a *FSCloudStorage) UpdateStream(path, resource string) (string, error) {
	a.init()
	path = a.ApplyPathPrefix(path)
	err := a.Adapter().UpdateStream(path, resource)
	return path, err
}

func (a *FSCloudStorage) MimeType(path string) (string, error) {
	a.init()
	path = a.ApplyPathPrefix(path)
	return a.Adapter().MimeType(path)
}

func (a *FSCloudStorage) Move(source, destination string) (bool, error) {
	a.init()
	source = a.ApplyPathPrefix(source)
	destination = a.ApplyPathPrefix(destination)
	return a.Adapter().Move(source, destination)
}

func (a *FSCloudStorage) Copy(source, destination string) (bool, error) {
	a.init()
	source = a.ApplyPathPrefix(source)
	destination = a.ApplyPathPrefix(destination)
	return a.Adapter().Copy(source, destination)
}

func (a *FSCloudStorage) DiskName() string {
	return flysystem.DiskNameGoogleCloudCloudStorage
}
