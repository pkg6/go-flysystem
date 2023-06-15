package fsoss

import (
	"bytes"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/pkg6/go-flysystem"
	"io"
	"net/http"
	"strconv"
	"sync"
)

type Config struct {
	Bucket          string `json:"bucket"`
	Endpoint        string `json:"endpoint"`
	AccessKeyID     string `json:"access_key_id"`
	AccessKeySecret string `json:"access_key_secret"`
	PathPrefix      string `json:"path_prefix"`
	OssConfig       *oss.Config
}
type FsOss struct {
	flysystem.AbstractAdapter
	Config Config
	Oss    *oss.Client
	lock   *sync.Mutex
}

func New(config Config) flysystem.IAdapter {
	return FsOss{Config: config}.Clone()
}

func (f FsOss) DiskName() string {
	return "oss"
}
func (f FsOss) Clone() flysystem.IAdapter {
	var err error
	if f.Config.Endpoint == "" {
		f.Config.Endpoint = "oss-cn-hangzhou.aliyuncs.com"
	}
	if f.Config.PathPrefix != "" {
		f.SetPathPrefix(f.Config.PathPrefix)
	}
	f.Oss, err = oss.New(f.Config.Endpoint, f.Config.AccessKeyID, f.Config.AccessKeySecret, func(client *oss.Client) {
		if f.Config.OssConfig != nil {
			client.Config = f.Config.OssConfig
		}
	})
	f.lock = &sync.Mutex{}
	if err != nil {
		panic(err)
	}
	return &f
}
func (f FsOss) Exists(path string) (bool, error) {
	f.lock.Lock()
	defer f.lock.Unlock()
	path = f.ApplyPathPrefix(path)
	bucket, err := f.Oss.Bucket(f.Config.Bucket)
	if err != nil {
		return false, err
	}
	return bucket.IsObjectExist(path)
}
func (f FsOss) WriteReader(path string, reader io.Reader) (string, error) {
	f.lock.Lock()
	defer f.lock.Unlock()
	path = f.ApplyPathPrefix(path)
	bucket, err := f.Oss.Bucket(f.Config.Bucket)
	if err != nil {
		return "", err
	}
	if err = bucket.PutObject(path, reader); err != nil {
		return "", err
	}
	return path, nil
}

func (f FsOss) Write(path string, contents []byte) (string, error) {
	return f.WriteReader(path, bytes.NewReader(contents))
}

func (f FsOss) WriteStream(path, resource string) (string, error) {
	f.lock.Lock()
	defer f.lock.Unlock()
	path = f.ApplyPathPrefix(path)
	bucket, err := f.Oss.Bucket(f.Config.Bucket)
	if err != nil {
		return "", err
	}
	if err = bucket.PutObjectFromFile(path, resource); err != nil {
		return "", err
	}
	return path, nil
}
func (f FsOss) Update(path string, contents []byte) (string, error) {
	return f.Write(path, contents)
}

func (f FsOss) UpdateStream(path, resource string) (string, error) {
	return f.WriteStream(path, resource)
}
func (f FsOss) Read(path string) ([]byte, error) {
	f.lock.Lock()
	defer f.lock.Unlock()
	path = f.ApplyPathPrefix(path)
	bucket, err := f.Oss.Bucket(f.Config.Bucket)
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

func (f FsOss) Delete(path string) (int64, error) {
	f.lock.Lock()
	defer f.lock.Unlock()
	path = f.ApplyPathPrefix(path)
	bucket, err := f.Oss.Bucket(f.Config.Bucket)
	if err != nil {
		return 0, err
	}
	if err = bucket.DeleteObject(path); err != nil {
		return 0, err
	}
	return 1, nil
}

func (f FsOss) DeleteDirectory(dirname string) (int64, error) {
	f.lock.Lock()
	defer f.lock.Unlock()
	dirname = f.ApplyPathPrefix(dirname)
	bucket, err := f.Oss.Bucket(f.Config.Bucket)
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
func (f FsOss) CreateDirectory(dirname string) error {
	f.lock.Lock()
	defer f.lock.Unlock()
	_, err := f.Write(dirname, []byte(""))
	return err
}

func (f FsOss) MimeType(path string) (string, error) {
	f.lock.Lock()
	defer f.lock.Unlock()
	meta, err := f.getObjectMeta(path)
	if err != nil {
		return "", err
	}
	return meta.Get("content-type"), nil
}

func (f FsOss) Size(path string) (int64, error) {
	f.lock.Lock()
	defer f.lock.Unlock()
	meta, err := f.getObjectMeta(path)
	if err != nil {
		return 0, err
	}
	i, err := strconv.ParseInt(meta.Get("content-length"), 10, 64)
	if err != nil {
		return 0, err
	}
	return i, nil
}
func (f FsOss) Move(source, destination string) (bool, error) {
	return f.copyObject(source, destination, true)
}

func (f FsOss) Copy(source, destination string) (bool, error) {
	return f.copyObject(source, destination, false)
}
func (f FsOss) copyObject(srcObjectKey, destObjectKey string, isDelete bool) (bool, error) {
	f.lock.Lock()
	defer f.lock.Unlock()
	srcObjectKey = f.ApplyPathPrefix(srcObjectKey)
	destObjectKey = f.ApplyPathPrefix(destObjectKey)
	bucket, err := f.Oss.Bucket(f.Config.Bucket)
	if err != nil {
		return false, err
	}
	_, err = bucket.CopyObject(srcObjectKey, destObjectKey)
	if err != nil {
		return false, err
	}
	if isDelete {
		err = bucket.DeleteObject(srcObjectKey)
	}
	return true, nil
}

func (f FsOss) getObjectMeta(path string) (header http.Header, err error) {
	path = f.ApplyPathPrefix(path)
	bucket, err := f.Oss.Bucket(f.Config.Bucket)
	if err != nil {
		return header, err
	}
	return bucket.GetObjectMeta(path)
}
