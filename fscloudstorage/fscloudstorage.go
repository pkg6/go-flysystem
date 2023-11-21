package fscloudstorage

import (
	"io"
	"net/url"
	"sync"
	"time"

	"github.com/pkg6/go-flysystem"
	"github.com/zzqqw/gfs"
	fscloudstorage2 "github.com/zzqqw/gfs/fscloudstorage"
	"google.golang.org/api/option"
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

func (a *FSCloudStorage) GFSAdapter() gfs.IAdapter {
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
	return a.GFSAdapter().URL(path)
}
func (a *FSCloudStorage) Exists(path string) (bool, error) {
	a.init()
	path = a.ApplyPathPrefix(path)
	return a.GFSAdapter().Exist(path)
}

func (a *FSCloudStorage) WriteReader(path string, reader io.Reader) (string, error) {
	a.init()
	path = a.ApplyPathPrefix(path)
	err := a.GFSAdapter().WriteReader(path, reader)
	return path, err
}

func (a *FSCloudStorage) Write(path string, contents []byte) (string, error) {
	a.init()
	path = a.ApplyPathPrefix(path)
	err := a.GFSAdapter().Write(path, contents)
	return path, err
}

func (a *FSCloudStorage) WriteStream(path, resource string) (string, error) {
	a.init()
	path = a.ApplyPathPrefix(path)
	err := a.GFSAdapter().WriteStream(path, resource)
	return path, err
}

func (a *FSCloudStorage) Read(path string) ([]byte, error) {
	a.init()
	path = a.ApplyPathPrefix(path)
	return a.GFSAdapter().Read(path)
}

func (a *FSCloudStorage) Delete(path string) (int64, error) {
	a.init()
	path = a.ApplyPathPrefix(path)
	return a.GFSAdapter().Delete(path)
}

func (a *FSCloudStorage) Size(path string) (int64, error) {
	a.init()
	path = a.ApplyPathPrefix(path)
	return a.GFSAdapter().Size(path)
}

func (a *FSCloudStorage) Update(path string, contents []byte) (string, error) {
	a.init()
	path = a.ApplyPathPrefix(path)
	err := a.GFSAdapter().Update(path, contents)
	return path, err
}

func (a *FSCloudStorage) UpdateStream(path, resource string) (string, error) {
	a.init()
	path = a.ApplyPathPrefix(path)
	err := a.GFSAdapter().UpdateStream(path, resource)
	return path, err
}

func (a *FSCloudStorage) MimeType(path string) (string, error) {
	a.init()
	path = a.ApplyPathPrefix(path)
	return a.GFSAdapter().MimeType(path)
}

func (a *FSCloudStorage) Move(source, destination string) (bool, error) {
	a.init()
	source = a.ApplyPathPrefix(source)
	destination = a.ApplyPathPrefix(destination)
	return a.GFSAdapter().Move(source, destination)
}

func (a *FSCloudStorage) Copy(source, destination string) (bool, error) {
	a.init()
	source = a.ApplyPathPrefix(source)
	destination = a.ApplyPathPrefix(destination)
	return a.GFSAdapter().Copy(source, destination)
}

func (a *FSCloudStorage) DiskName() string {
	return flysystem.DiskNameGoogleCloudCloudStorage
}
