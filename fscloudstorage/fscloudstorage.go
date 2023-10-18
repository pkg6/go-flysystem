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
func (f *FSCloudStorage) URL(path string) (*url.URL, error) {
	path = f.ApplyPathPrefix(path)
	return f.Adapter().URL(path)
}
func (f *FSCloudStorage) Adapter() *fscloudstorage2.Adapter {
	if f.lock == nil {
		f.lock = &sync.Mutex{}
	}
	f.SetPathPrefix(f.Config.PathPrefix)
	return fscloudstorage2.NewGCS(&fscloudstorage2.Config{
		CDN:             f.Config.CDN,
		Bucket:          f.Config.Bucket,
		WithTimeout:     f.Config.WithTimeout,
		CredentialsFile: f.Config.CredentialsFile,
		Option:          f.Config.Option,
	})
}

func (f *FSCloudStorage) Exists(path string) (bool, error) {
	path = f.ApplyPathPrefix(path)
	return f.Adapter().Exist(path)
}

func (f *FSCloudStorage) WriteReader(path string, reader io.Reader) (string, error) {
	path = f.ApplyPathPrefix(path)
	err := f.Adapter().WriteReader(path, reader)
	return path, err
}

func (f *FSCloudStorage) Write(path string, contents []byte) (string, error) {
	path = f.ApplyPathPrefix(path)
	err := f.Adapter().Write(path, contents)
	return path, err
}

func (f *FSCloudStorage) WriteStream(path, resource string) (string, error) {
	path = f.ApplyPathPrefix(path)
	err := f.Adapter().WriteStream(path, resource)
	return path, err
}

func (f *FSCloudStorage) Read(path string) ([]byte, error) {
	path = f.ApplyPathPrefix(path)
	return f.Adapter().Read(path)
}

func (f *FSCloudStorage) Delete(path string) (int64, error) {
	path = f.ApplyPathPrefix(path)
	return f.Adapter().Delete(path)
}

func (f *FSCloudStorage) Size(path string) (int64, error) {
	path = f.ApplyPathPrefix(path)
	return f.Adapter().Size(path)
}

func (f *FSCloudStorage) Update(path string, contents []byte) (string, error) {
	path = f.ApplyPathPrefix(path)
	err := f.Adapter().Update(path, contents)
	return path, err
}

func (f *FSCloudStorage) UpdateStream(path, resource string) (string, error) {
	path = f.ApplyPathPrefix(path)
	err := f.Adapter().UpdateStream(path, resource)
	return path, err
}

func (f *FSCloudStorage) MimeType(path string) (string, error) {
	path = f.ApplyPathPrefix(path)
	return f.Adapter().MimeType(path)
}

func (f *FSCloudStorage) Move(source, destination string) (bool, error) {
	source = f.ApplyPathPrefix(source)
	destination = f.ApplyPathPrefix(destination)
	return f.Adapter().Move(source, destination)
}

func (f *FSCloudStorage) Copy(source, destination string) (bool, error) {
	source = f.ApplyPathPrefix(source)
	destination = f.ApplyPathPrefix(destination)
	return f.Adapter().Copy(source, destination)
}

func (f *FSCloudStorage) DiskName() string {
	return flysystem.DiskNameGoogleCloudCloudStorage
}
