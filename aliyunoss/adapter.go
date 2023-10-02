package aliyunoss

import (
	"bytes"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/pkg6/go-flysystem"
	"io"
	"net/http"
	"strconv"
	"sync"
)

var (
	DefaultEndpoint = "oss-cn-hangzhou.aliyuncs.com"
)

type Config struct {
	Bucket          string `json:"bucket"`
	Endpoint        string `json:"endpoint"`
	AccessKeyID     string `json:"access_key_id"`
	AccessKeySecret string `json:"access_key_secret"`
	PathPrefix      string `json:"path_prefix"`
	OssConfig       *oss.Config
}

type Adapter struct {
	Config *Config
	lock   *sync.Mutex
}

func New(config *Config) flysystem.IAdapter {
	a := Adapter{Config: config}
	return a.Clone()
}

func (a *Adapter) DiskName() string {
	return flysystem.DiskNameOSS
}

func (a *Adapter) OSSClient() (*oss.Client, error) {
	return oss.New(a.Config.Endpoint, a.Config.AccessKeyID, a.Config.AccessKeySecret, func(client *oss.Client) {
		if a.Config.OssConfig != nil {
			client.Config = a.Config.OssConfig
		}
	})
}

func (a *Adapter) OSSBucket() (*oss.Bucket, error) {
	client, err := a.OSSClient()
	if err != nil {
		return nil, err
	}
	return client.Bucket(a.Config.Bucket)
}

func (a *Adapter) Clone() flysystem.IAdapter {
	if a.Config.Endpoint == "" {
		a.Config.Endpoint = DefaultEndpoint
	}
	a.lock = &sync.Mutex{}
	return a
}

func (a *Adapter) Exists(path string) (bool, error) {
	a.lock.Lock()
	defer a.lock.Unlock()
	bucket, err := a.OSSBucket()
	if err != nil {
		return false, err
	}
	return bucket.IsObjectExist(path)
}
func (a *Adapter) WriteReader(path string, reader io.Reader) (string, error) {
	a.lock.Lock()
	defer a.lock.Unlock()
	bucket, err := a.OSSBucket()
	if err != nil {
		return "", err
	}
	if err = bucket.PutObject(path, reader); err != nil {
		return "", err
	}
	return path, nil
}

func (a *Adapter) Write(path string, contents []byte) (string, error) {
	return a.WriteReader(path, bytes.NewReader(contents))
}

func (a *Adapter) WriteStream(path, resource string) (string, error) {
	a.lock.Lock()
	defer a.lock.Unlock()
	bucket, err := a.OSSBucket()
	if err != nil {
		return "", err
	}
	if err = bucket.PutObjectFromFile(path, resource); err != nil {
		return "", err
	}
	return path, nil
}
func (a *Adapter) Update(path string, contents []byte) (string, error) {
	return a.Write(path, contents)
}

func (a *Adapter) UpdateStream(path, resource string) (string, error) {
	return a.WriteStream(path, resource)
}
func (a *Adapter) Read(path string) ([]byte, error) {
	a.lock.Lock()
	defer a.lock.Unlock()
	bucket, err := a.OSSBucket()
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
	bucket, err := a.OSSBucket()
	if err != nil {
		return 0, err
	}
	if err = bucket.DeleteObject(path); err != nil {
		return 0, err
	}
	return 1, nil
}

func (a *Adapter) DeleteDirectory(dirname string) (int64, error) {
	a.lock.Lock()
	defer a.lock.Unlock()
	bucket, err := a.OSSBucket()
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
func (a *Adapter) CreateDirectory(dirname string) error {
	a.lock.Lock()
	defer a.lock.Unlock()
	_, err := a.Write(dirname, []byte(""))
	return err
}

func (a *Adapter) MimeType(path string) (string, error) {
	a.lock.Lock()
	defer a.lock.Unlock()
	meta, err := a.getObjectMeta(path)
	if err != nil {
		return "", err
	}
	return meta.Get("content-type"), nil
}

func (a *Adapter) Size(path string) (int64, error) {
	a.lock.Lock()
	defer a.lock.Unlock()
	meta, err := a.getObjectMeta(path)
	if err != nil {
		return 0, err
	}
	i, err := strconv.ParseInt(meta.Get("content-length"), 10, 64)
	if err != nil {
		return 0, err
	}
	return i, nil
}
func (a *Adapter) Move(source, destination string) (bool, error) {
	return a.copyObject(source, destination, true)
}

func (a *Adapter) Copy(source, destination string) (bool, error) {
	return a.copyObject(source, destination, false)
}

func (a *Adapter) copyObject(srcObjectKey, destObjectKey string, isDelete bool) (bool, error) {
	a.lock.Lock()
	defer a.lock.Unlock()
	bucket, err := a.OSSBucket()
	if err != nil {
		return false, err
	}
	_, err = bucket.CopyObject(srcObjectKey, destObjectKey)
	if err != nil {
		return false, err
	}
	if isDelete {
		_ = bucket.DeleteObject(srcObjectKey)
	}
	return true, nil
}
func (a *Adapter) getObjectMeta(path string) (header http.Header, err error) {
	bucket, err := a.OSSBucket()
	if err != nil {
		return header, err
	}
	return bucket.GetObjectMeta(path)
}
