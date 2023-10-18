package fsoss

import (
	"bytes"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	v2 "github.com/pkg6/go-flysystem/v2"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"sync"
)

type Adapter struct {
	Config *Config
	lock   *sync.Mutex
}

func New(config v2.IAdapterConfig) v2.IAdapter {
	return config.New()
}

func NewOSS(config *Config) *Adapter {
	a := &Adapter{Config: config}
	if a.Config.Endpoint == "" {
		a.Config.Endpoint = Endpoint(RegionCnHangzhou)
	}
	a.lock = &sync.Mutex{}
	return a
}

func (a *Adapter) DiskName() string {
	return v2.DiskNameOSS
}

func (a *Adapter) Client() (*oss.Client, error) {
	return oss.New(a.Config.Endpoint, a.Config.AccessKeyID, a.Config.AccessKeySecret, func(client *oss.Client) {
		if a.Config.OssConfig != nil {
			client.Config = a.Config.OssConfig
		}
	})
}

func (a *Adapter) Bucket() (*oss.Bucket, error) {
	client, err := a.Client()
	if err != nil {
		return nil, err
	}
	return client.Bucket(a.Config.Bucket)
}
func (a *Adapter) CopyObject(srcObjectKey, destObjectKey string, isDelete bool) (bool, error) {
	a.lock.Lock()
	defer a.lock.Unlock()
	bucket, err := a.Bucket()
	if err != nil {
		return false, err
	}
	_, err = bucket.CopyObject(srcObjectKey, destObjectKey)
	if err != nil {
		return false, err
	}
	if isDelete {
		defer func() {
			_ = bucket.DeleteObject(srcObjectKey)
		}()
	}
	return true, nil
}
func (a *Adapter) Meta(path string) (header http.Header, err error) {
	bucket, err := a.Bucket()
	if err != nil {
		return header, err
	}
	return bucket.GetObjectMeta(path)
}
func (a *Adapter) URL(path string) (*url.URL, error) {
	return a.Config.URL(path)
}
func (a *Adapter) Exist(path string) (bool, error) {
	a.lock.Lock()
	defer a.lock.Unlock()
	bucket, err := a.Bucket()
	if err != nil {
		return false, err
	}
	return bucket.IsObjectExist(path)
}
func (a *Adapter) WriteReader(path string, reader io.Reader) error {
	a.lock.Lock()
	defer a.lock.Unlock()
	bucket, err := a.Bucket()
	if err != nil {
		return err
	}
	return bucket.PutObject(path, reader)
}

func (a *Adapter) Write(path string, contents []byte) error {
	return a.WriteReader(path, bytes.NewReader(contents))
}

func (a *Adapter) WriteStream(path, resource string) error {
	a.lock.Lock()
	defer a.lock.Unlock()
	bucket, err := a.Bucket()
	if err != nil {
		return err
	}
	return bucket.PutObjectFromFile(path, resource)
}
func (a *Adapter) Update(path string, contents []byte) error {
	return a.Write(path, contents)
}

func (a *Adapter) UpdateStream(path, resource string) error {
	return a.WriteStream(path, resource)
}
func (a *Adapter) Read(path string) ([]byte, error) {
	a.lock.Lock()
	defer a.lock.Unlock()
	bucket, err := a.Bucket()
	if err != nil {
		return nil, err
	}
	object, err := bucket.GetObject(path)
	if err != nil {
		return nil, err
	}
	defer object.Close()
	contents, err := io.ReadAll(object)
	if err != nil {
		return nil, err
	}
	return contents, err
}

func (a *Adapter) Delete(path string) (int64, error) {
	a.lock.Lock()
	defer a.lock.Unlock()
	bucket, err := a.Bucket()
	if err != nil {
		return 0, err
	}
	if err = bucket.DeleteObject(path); err != nil {
		return 0, err
	}
	return 1, nil
}

func (a *Adapter) MimeType(path string) (string, error) {
	a.lock.Lock()
	defer a.lock.Unlock()
	meta, err := a.Meta(path)
	if err != nil {
		return "", err
	}
	return meta.Get(v2.HeaderGetContentType), nil
}

func (a *Adapter) Size(path string) (int64, error) {
	a.lock.Lock()
	defer a.lock.Unlock()
	meta, err := a.Meta(path)
	if err != nil {
		return 0, err
	}
	i, err := strconv.ParseInt(meta.Get(v2.HeaderGetLength), 10, 64)
	if err != nil {
		return 0, err
	}
	return i, nil
}
func (a *Adapter) Move(source, destination string) (bool, error) {
	return a.CopyObject(source, destination, true)
}

func (a *Adapter) Copy(source, destination string) (bool, error) {
	return a.CopyObject(source, destination, false)
}
