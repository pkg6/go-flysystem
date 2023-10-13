package fscos

import (
	"github.com/pkg6/go-flysystem"
	"github.com/pkg6/go-flysystem/v2"
	fscos2 "github.com/pkg6/go-flysystem/v2/fscos"
	"io"
)

type Config struct {
	// 存储桶名称，由 bucketname-appid 组成，appid 必须填入，可以在 COS 控制台查看存储桶名称。 https://console.cloud.tencent.com/cos5/bucket
	// 替换为用户的 region，存储桶 region 可以在 COS 控制台“存储桶概览”查看 https://console.cloud.tencent.com/
	BucketURL string
	// 用户的 SecretId，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参见 https://cloud.tencent.com/document/product/598/37140
	// 用户的 SecretKey，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参见 https://cloud.tencent.com/document/product/598/37140
	SecretID, SecretKey string
	PathPrefix          string
}
type FSCos struct {
	v2.AbstractAdapter
	Config *Config
}

func New(config *Config) flysystem.IAdapter {
	return &FSCos{Config: config}
}

func (f *FSCos) Adapter() *fscos2.Adapter {
	return fscos2.NewCOS(&fscos2.Config{
		BucketURL: f.Config.BucketURL,
		SecretID:  f.Config.SecretID,
		SecretKey: f.Config.SecretKey,
	})
}

func (f *FSCos) Exists(path string) (bool, error) {
	path = f.ApplyPathPrefix(path)
	return f.Adapter().Exist(path)
}

func (f *FSCos) WriteReader(path string, reader io.Reader) (string, error) {
	path = f.ApplyPathPrefix(path)
	err := f.Adapter().WriteReader(path, reader)
	return path, err
}

func (f *FSCos) Write(path string, contents []byte) (string, error) {
	path = f.ApplyPathPrefix(path)
	err := f.Adapter().Write(path, contents)
	return path, err
}

func (f *FSCos) WriteStream(path, resource string) (string, error) {
	path = f.ApplyPathPrefix(path)
	err := f.Adapter().WriteStream(path, resource)
	return path, err
}

func (f *FSCos) Read(path string) ([]byte, error) {
	path = f.ApplyPathPrefix(path)
	return f.Adapter().Read(path)
}

func (f *FSCos) Delete(path string) (int64, error) {
	path = f.ApplyPathPrefix(path)
	return f.Adapter().Delete(path)
}

func (f *FSCos) Size(path string) (int64, error) {
	path = f.ApplyPathPrefix(path)
	return f.Adapter().Size(path)
}

func (f *FSCos) Update(path string, contents []byte) (string, error) {
	path = f.ApplyPathPrefix(path)
	err := f.Adapter().Update(path, contents)
	return path, err
}

func (f *FSCos) UpdateStream(path, resource string) (string, error) {
	path = f.ApplyPathPrefix(path)
	err := f.Adapter().UpdateStream(path, resource)
	return path, err
}

func (f *FSCos) MimeType(path string) (string, error) {
	path = f.ApplyPathPrefix(path)
	return f.Adapter().MimeType(path)
}

func (f *FSCos) Move(source, destination string) (bool, error) {
	source = f.ApplyPathPrefix(source)
	destination = f.ApplyPathPrefix(destination)
	return f.Adapter().Move(source, destination)
}

func (f *FSCos) Copy(source, destination string) (bool, error) {
	source = f.ApplyPathPrefix(source)
	destination = f.ApplyPathPrefix(destination)
	return f.Adapter().Copy(source, destination)
}

func (f *FSCos) DiskName() string {
	return flysystem.DiskNameCOS
}
