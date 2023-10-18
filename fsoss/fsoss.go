package fsoss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/pkg6/go-flysystem"
	"github.com/pkg6/go-flysystem/gfs"
	fsoss2 "github.com/pkg6/go-flysystem/gfs/fsoss"
	"io"
	"net/url"
	"sync"
)

var (
	DefaultEndpoint = "oss-cn-hangzhou.aliyuncs.com"
)

type Config struct {
	CDN             string
	Bucket          string
	Endpoint        string
	AccessKeyID     string
	AccessKeySecret string
	OssConfig       *oss.Config
	PathPrefix      string
}
type FsOss struct {
	gfs.AbstractAdapter
	Config *Config
	lock   *sync.Mutex
}

func New(config *Config) flysystem.IAdapter {
	f := &FsOss{Config: config}
	return f
}

func (f *FsOss) Adapter() *fsoss2.Adapter {
	if f.Config.Endpoint == "" {
		f.Config.Endpoint = DefaultEndpoint
	}
	if f.lock == nil {
		f.lock = &sync.Mutex{}
	}
	f.SetPathPrefix(f.Config.PathPrefix)
	return fsoss2.NewOSS(&fsoss2.Config{
		CDN:             f.Config.CDN,
		Bucket:          f.Config.Bucket,
		Endpoint:        f.Config.Endpoint,
		AccessKeyID:     f.Config.AccessKeyID,
		AccessKeySecret: f.Config.AccessKeySecret,
		OssConfig:       f.Config.OssConfig,
	})
}

func (f *FsOss) DiskName() string {
	return flysystem.DiskNameOSS
}

func (f *FsOss) Exists(path string) (bool, error) {
	path = f.ApplyPathPrefix(path)
	return f.Adapter().Exist(path)
}
func (f *FsOss) WriteReader(path string, reader io.Reader) (string, error) {
	path = f.ApplyPathPrefix(path)
	err := f.Adapter().WriteReader(path, reader)
	return path, err
}

func (f *FsOss) Write(path string, contents []byte) (string, error) {
	path = f.ApplyPathPrefix(path)
	err := f.Adapter().Write(path, contents)
	return path, err
}

func (f *FsOss) WriteStream(path, resource string) (string, error) {
	path = f.ApplyPathPrefix(path)
	err := f.Adapter().WriteStream(path, resource)
	return path, err
}
func (f *FsOss) Update(path string, contents []byte) (string, error) {
	path = f.ApplyPathPrefix(path)
	err := f.Adapter().Update(path, contents)
	return path, err
}
func (f *FsOss) URL(path string) (*url.URL, error) {
	path = f.ApplyPathPrefix(path)
	return f.Adapter().URL(path)
}

func (f *FsOss) UpdateStream(path, resource string) (string, error) {
	path = f.ApplyPathPrefix(path)
	err := f.Adapter().UpdateStream(path, resource)
	return path, err
}
func (f *FsOss) Read(path string) ([]byte, error) {
	path = f.ApplyPathPrefix(path)
	return f.Adapter().Read(path)
}

func (f *FsOss) Delete(path string) (int64, error) {
	path = f.ApplyPathPrefix(path)
	return f.Adapter().Delete(path)
}

func (f *FsOss) MimeType(path string) (string, error) {
	path = f.ApplyPathPrefix(path)
	return f.Adapter().MimeType(path)
}

func (f *FsOss) Size(path string) (int64, error) {
	path = f.ApplyPathPrefix(path)
	return f.Adapter().Size(path)
}
func (f *FsOss) Move(source, destination string) (bool, error) {
	source = f.ApplyPathPrefix(source)
	destination = f.ApplyPathPrefix(destination)
	return f.Adapter().Move(source, destination)
}

func (f *FsOss) Copy(source, destination string) (bool, error) {
	source = f.ApplyPathPrefix(source)
	destination = f.ApplyPathPrefix(destination)
	return f.Adapter().Copy(source, destination)
}

func (f *FsOss) DeleteDirectory(dirname string) (int64, error) {
	f.lock.Lock()
	defer f.lock.Unlock()
	client, err := f.Adapter().Client()
	if err != nil {
		return 0, err
	}
	dirname = f.ApplyPathPrefix(dirname)
	bucket, err := client.Bucket(f.Config.Bucket)
	if err != nil {
		return 0, err
	}
	marker := oss.Marker("")
	prefix := oss.Prefix(dirname)
	var count int64
	for {
		lor, err := bucket.ListObjects(marker, prefix)
		if err != nil {
			return 0, err
		}
		var objects []string
		for _, object := range lor.Objects {
			objects = append(objects, object.Key)
		}
		delRes, err := bucket.DeleteObjects(objects, oss.DeleteObjectsQuiet(true))
		if err != nil {
			return 0, err
		}
		if len(delRes.DeletedObjects) > 0 {
			return 0, err
		}
		count += int64(len(objects))
		prefix = oss.Prefix(lor.Prefix)
		marker = oss.Marker(lor.NextMarker)
		if !lor.IsTruncated {
			break
		}
	}
	return count, nil
}
func (f *FsOss) CreateDirectory(dirname string) error {
	f.lock.Lock()
	defer f.lock.Unlock()
	_, err := f.Write(dirname, []byte(""))
	return err
}
