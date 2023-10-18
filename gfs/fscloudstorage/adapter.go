package fscloudstorage

import (
	"bytes"
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"github.com/pkg6/go-flysystem/gfs"
	"google.golang.org/api/option"
	"io"
	"net/url"
	"os"
	"sync"
)

// Adapter
// api key https://console.cloud.google.com/apis/dashboard
//
// CredentialsFile
// 1. IAM和管理->服务账号->创建账号->密钥->添加密钥->密钥类型json->创建
// 2. Cloud Storage->创建->指定存储桶的名称【test】->点击桶列表中【test】->权限->授予访问权限->添加主账号（allUsers）->分配角色（Storage Admin）
type Adapter struct {
	Config             *Config
	ctx                context.Context
	lock               *sync.Mutex
	closeTimeoutCancel []func()
	closeClient        *storage.Client
}

func New(config gfs.IAdapterConfig) gfs.IAdapter {
	return config.New()
}
func NewGCS(config *Config) *Adapter {
	if config.CredentialsFile != "" {
		config.Option = append(config.Option, option.WithCredentialsFile(config.CredentialsFile))
	}
	a := &Adapter{Config: config}
	if a.Config.WithTimeout == 0 {
		a.Config.WithTimeout = DefaultWithTimeout
	}
	a.ctx = context.Background()
	a.lock = &sync.Mutex{}
	return a
}

func (a *Adapter) DiskName() string {
	return gfs.DiskNameGoogleCloudCloudStorage
}

func (a *Adapter) Client() (*storage.Client, error) {
	client, err := storage.NewClient(a.ctx, a.Config.Option...)
	if err != nil {
		return nil, err
	}
	return client, err
}

func (a *Adapter) StorageObject(object string) (*storage.ObjectHandle, error) {
	client, err := a.Client()
	if err != nil {
		return nil, err
	}
	obj := client.Bucket(a.Config.Bucket).Object(object).If(
		storage.Conditions{DoesNotExist: true},
	)
	return obj, nil
}

func (a *Adapter) StorageObjectTimeout(object string) (*storage.ObjectHandle, context.Context, error) {
	objectHandle, err := a.StorageObject(object)
	if err != nil {
		return nil, a.ctx, err
	}
	ctx, cancel := context.WithTimeout(a.ctx, a.Config.WithTimeout)
	a.closeTimeoutCancel = append(a.closeTimeoutCancel, cancel)
	return objectHandle, ctx, nil
}
func (a *Adapter) GetMetadata(object string) (*storage.ObjectAttrs, error) {
	obj, ctx, err := a.StorageObjectTimeout(object)
	defer a.Close()
	if err != nil {
		return nil, err
	}
	return obj.Attrs(ctx)
}
func (a *Adapter) CopyObject(source, destination string, deleteSource bool) (bool, error) {
	src, ctx, err := a.StorageObjectTimeout(source)
	defer a.Close()
	if err != nil {
		return false, err
	}
	dst, _, err := a.StorageObjectTimeout(source)
	defer a.Close()
	if err != nil {
		return false, err
	}
	if _, err := dst.CopierFrom(src).Run(ctx); err != nil {
		return false, fmt.Errorf("Object(%q).CopierFrom(%q).Run: %w", source, destination, err)
	}
	if deleteSource {
		defer func() {
			_ = src.Delete(ctx)
		}()
	}
	return true, nil
}
func (a *Adapter) URL(path string) (*url.URL, error) {
	return a.Config.URL(path)
}

func (a *Adapter) Exist(path string) (bool, error) {
	size, err := a.Size(path)
	if size > 0 && err == nil {
		return true, nil
	}
	return false, err
}

func (a *Adapter) WriteReader(path string, reader io.Reader) error {
	obj, ctx, err := a.StorageObjectTimeout(path)
	defer a.Close()
	if err != nil {
		return fmt.Errorf("storage.storageObject: %w", err)
	}
	// Upload an object with storage.Writer.
	wc := obj.NewWriter(ctx)
	wc.ChunkSize = 0 // note retries are not supported for chunk size 0.
	if _, err = io.Copy(wc, reader); err != nil {
		return fmt.Errorf("io.Copy: %w", err)
	}
	// Data can continue to be added to the file until the writer is closed.
	if err := wc.Close(); err != nil {
		return fmt.Errorf("Writer.Close: %w", err)
	}
	return nil
}

func (a *Adapter) Write(path string, contents []byte) error {
	return a.WriteReader(path, bytes.NewBuffer(contents))
}

func (a *Adapter) WriteStream(path, resource string) error {
	f, err := os.Open(resource)
	if err != nil {
		return fmt.Errorf("os.Open: %w", err)
	}
	return a.WriteReader(path, f)
}

func (a *Adapter) Read(path string) ([]byte, error) {
	obj, ctx, err := a.StorageObjectTimeout(path)
	defer a.Close()
	if err != nil {
		return nil, err
	}
	rc, err := obj.NewReader(ctx)
	if err != nil {
		return nil, fmt.Errorf("Object(%q).NewReader: %w", path, err)
	}
	defer rc.Close()
	data, err := io.ReadAll(rc)
	if err != nil {
		return nil, fmt.Errorf("io.ReadAll: %w", err)
	}
	return data, nil
}

func (a *Adapter) Delete(path string) (int64, error) {
	obj, ctx, err := a.StorageObjectTimeout(path)
	defer a.Close()
	if err != nil {
		return 0, err
	}
	err = obj.Delete(ctx)
	if err != nil {
		return 0, err
	}
	return 1, nil
}

func (a *Adapter) Size(path string) (int64, error) {
	metadata, err := a.GetMetadata(path)
	if err != nil {
		return 0, err
	}
	return metadata.Size, nil
}

func (a *Adapter) Update(path string, contents []byte) error {
	return a.Write(path, contents)
}

func (a *Adapter) UpdateStream(path, resource string) error {
	return a.WriteStream(path, resource)
}

func (a *Adapter) MimeType(path string) (string, error) {
	metadata, err := a.GetMetadata(path)
	if err != nil {
		return "", err
	}
	return metadata.ContentType, nil
}

func (a *Adapter) Move(source, destination string) (bool, error) {
	return a.CopyObject(source, destination, true)
}

func (a *Adapter) Copy(source, destination string) (bool, error) {
	return a.CopyObject(source, destination, false)
}

func (a *Adapter) Close() {
	for _, f := range a.closeTimeoutCancel {
		f()
	}
	_ = a.closeClient.Close()
}
