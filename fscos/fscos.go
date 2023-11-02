package fscos

import (
	"io"
	"net/url"
	"sync"

	"github.com/pkg6/go-flysystem"
	"github.com/zzqqw/gfs"
	fscos2 "github.com/zzqqw/gfs/fscos"
)

type Config struct {
	CDN string
	// 存储桶名称，由 bucketname-appid 组成，appid 必须填入，可以在 COS 控制台查看存储桶名称。 https://console.cloud.tencent.com/cos5/bucket
	// 替换为用户的 region，存储桶 region 可以在 COS 控制台“存储桶概览”查看 https://console.cloud.tencent.com/
	BucketURL string
	// 用户的 SecretId，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参见 https://cloud.tencent.com/document/product/598/37140
	// 用户的 SecretKey，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参见 https://cloud.tencent.com/document/product/598/37140
	SecretID, SecretKey string
	PathPrefix          string
}
type FSCos struct {
	gfs.AbstractAdapter
	Config *Config
	lock   *sync.Mutex
}

func New(config *Config) flysystem.IAdapter {
	return &FSCos{Config: config}
}
func (a *FSCos) init() {
	if a.lock == nil {
		a.lock = &sync.Mutex{}
	}
	a.SetPathPrefix(a.Config.PathPrefix)
}
func (a *FSCos) Adapter() *fscos2.Adapter {
	return fscos2.NewCOS(&fscos2.Config{
		CDN:       a.Config.CDN,
		BucketURL: a.Config.BucketURL,
		SecretID:  a.Config.SecretID,
		SecretKey: a.Config.SecretKey,
	})
}
func (a *FSCos) URL(path string) (*url.URL, error) {
	a.init()
	path = a.ApplyPathPrefix(path)
	return a.Adapter().URL(path)
}
func (a *FSCos) Exists(path string) (bool, error) {
	a.init()
	path = a.ApplyPathPrefix(path)
	return a.Adapter().Exist(path)
}

func (a *FSCos) WriteReader(path string, reader io.Reader) (string, error) {
	a.init()
	path = a.ApplyPathPrefix(path)
	err := a.Adapter().WriteReader(path, reader)
	return path, err
}

func (a *FSCos) Write(path string, contents []byte) (string, error) {
	a.init()
	path = a.ApplyPathPrefix(path)
	err := a.Adapter().Write(path, contents)
	return path, err
}

func (a *FSCos) WriteStream(path, resource string) (string, error) {
	a.init()
	path = a.ApplyPathPrefix(path)
	err := a.Adapter().WriteStream(path, resource)
	return path, err
}

func (a *FSCos) Read(path string) ([]byte, error) {
	a.init()
	path = a.ApplyPathPrefix(path)
	return a.Adapter().Read(path)
}

func (a *FSCos) Delete(path string) (int64, error) {
	a.init()
	path = a.ApplyPathPrefix(path)
	return a.Adapter().Delete(path)
}

func (a *FSCos) Size(path string) (int64, error) {
	a.init()
	path = a.ApplyPathPrefix(path)
	return a.Adapter().Size(path)
}

func (a *FSCos) Update(path string, contents []byte) (string, error) {
	a.init()
	path = a.ApplyPathPrefix(path)
	err := a.Adapter().Update(path, contents)
	return path, err
}

func (a *FSCos) UpdateStream(path, resource string) (string, error) {
	a.init()
	path = a.ApplyPathPrefix(path)
	err := a.Adapter().UpdateStream(path, resource)
	return path, err
}

func (a *FSCos) MimeType(path string) (string, error) {
	a.init()
	path = a.ApplyPathPrefix(path)
	return a.Adapter().MimeType(path)
}

func (a *FSCos) Move(source, destination string) (bool, error) {
	a.init()
	source = a.ApplyPathPrefix(source)
	destination = a.ApplyPathPrefix(destination)
	return a.Adapter().Move(source, destination)
}

func (a *FSCos) Copy(source, destination string) (bool, error) {
	a.init()
	source = a.ApplyPathPrefix(source)
	destination = a.ApplyPathPrefix(destination)
	return a.Adapter().Copy(source, destination)
}

func (a *FSCos) DiskName() string {
	return flysystem.DiskNameCOS
}
