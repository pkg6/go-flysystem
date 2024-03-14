package fsoss

import (
	"github.com/pkg6/gfs/ossfs"
	"io"
	"net/url"
	"sync"

	"github.com/pkg6/gfs"
	"github.com/pkg6/go-flysystem"
)

var (
	DefaultEndpoint = "oss-cn-hangzhou.aliyuncs.com"
)

type FsOss struct {
	gfs.AbstractAdapter
	Config *Config
	lock   *sync.Mutex
}

func New(config *Config) flysystem.IAdapter {
	return NewOSS(config)
}
func NewOSS(config *Config) *FsOss {
	if config.Endpoint == "" {
		config.Endpoint = DefaultEndpoint
	}
	f := &FsOss{Config: config, lock: &sync.Mutex{}}
	f.SetPathPrefix(f.Config.PathPrefix)
	return f
}

func (f *FsOss) GFSAdapter() gfs.IAdapter {
	return ossfs.NewOSS(&ossfs.Config{
		CDN:             f.Config.CDN,
		Bucket:          f.Config.Bucket,
		Endpoint:        f.Config.Endpoint,
		AccessKeyID:     f.Config.AccessKeyID,
		AccessKeySecret: f.Config.AccessKeySecret,
		Config:          f.Config.Config,
	})
}

func (f *FsOss) DiskName() string {
	return flysystem.DiskNameOSS
}

func (f *FsOss) Exists(path string) (bool, error) {
	path = f.ApplyPathPrefix(path)
	return f.GFSAdapter().Exist(path)
}
func (f *FsOss) WriteReader(path string, reader io.Reader) (string, error) {
	path = f.ApplyPathPrefix(path)
	err := f.GFSAdapter().WriteReader(path, reader)
	return path, err
}

func (f *FsOss) Write(path string, contents []byte) (string, error) {

	path = f.ApplyPathPrefix(path)
	err := f.GFSAdapter().Write(path, contents)
	return path, err
}

func (f *FsOss) WriteStream(path, resource string) (string, error) {
	path = f.ApplyPathPrefix(path)
	err := f.GFSAdapter().WriteStream(path, resource)
	return path, err
}
func (f *FsOss) Update(path string, contents []byte) (string, error) {

	path = f.ApplyPathPrefix(path)
	err := f.GFSAdapter().Update(path, contents)
	return path, err
}
func (f *FsOss) URL(path string) (*url.URL, error) {

	path = f.ApplyPathPrefix(path)
	return f.GFSAdapter().URL(path)
}

func (f *FsOss) UpdateStream(path, resource string) (string, error) {

	path = f.ApplyPathPrefix(path)
	err := f.GFSAdapter().UpdateStream(path, resource)
	return path, err
}
func (f *FsOss) Read(path string) ([]byte, error) {

	path = f.ApplyPathPrefix(path)
	return f.GFSAdapter().Read(path)
}

func (f *FsOss) Delete(path string) (int64, error) {

	path = f.ApplyPathPrefix(path)
	return f.GFSAdapter().Delete(path)
}

func (f *FsOss) MimeType(path string) (string, error) {

	path = f.ApplyPathPrefix(path)
	return f.GFSAdapter().MimeType(path)
}

func (f *FsOss) Size(path string) (int64, error) {

	path = f.ApplyPathPrefix(path)
	return f.GFSAdapter().Size(path)
}
func (f *FsOss) Move(source, destination string) (bool, error) {

	source = f.ApplyPathPrefix(source)
	destination = f.ApplyPathPrefix(destination)
	return f.GFSAdapter().Move(source, destination)
}

func (f *FsOss) Copy(source, destination string) (bool, error) {

	source = f.ApplyPathPrefix(source)
	destination = f.ApplyPathPrefix(destination)
	return f.GFSAdapter().Copy(source, destination)
}

//
//func (f *FsOss) DeleteDirectory(dirname string) (int64, error) {
//	f.lock.Lock()
//	defer f.lock.Unlock()
//	client, err := f.GFSAdapter().Client()
//	if err != nil {
//		return 0, err
//	}
//	dirname = f.ApplyPathPrefix(dirname)
//	bucket, err := client.Bucket(f.Config.Bucket)
//	if err != nil {
//		return 0, err
//	}
//	marker := oss.Marker("")
//	prefix := oss.Prefix(dirname)
//	var count int64
//	for {
//		lor, err := bucket.ListObjects(marker, prefix)
//		if err != nil {
//			return 0, err
//		}
//		var objects []string
//		for _, object := range lor.Objects {
//			objects = append(objects, object.Key)
//		}
//		delRes, err := bucket.DeleteObjects(objects, oss.DeleteObjectsQuiet(true))
//		if err != nil {
//			return 0, err
//		}
//		if len(delRes.DeletedObjects) > 0 {
//			return 0, err
//		}
//		count += int64(len(objects))
//		prefix = oss.Prefix(lor.Prefix)
//		marker = oss.Marker(lor.NextMarker)
//		if !lor.IsTruncated {
//			break
//		}
//	}
//	return count, nil
//}

func (f *FsOss) CreateDirectory(dirname string) error {
	f.lock.Lock()
	defer f.lock.Unlock()
	_, err := f.Write(dirname, []byte(""))
	return err
}
