package fscos

import (
	"context"
	"fmt"
	v2 "github.com/pkg6/go-flysystem/v2"
	"github.com/tencentyun/cos-go-sdk-v5"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

type Adapter struct {
	Config *Config
	lock   *sync.Mutex
}

func (a *Adapter) URL(path string) (*url.URL, error) {
	return a.Config.URL(path)
}

func New(config v2.IAdapterConfig) v2.IAdapter {
	return config.New()
}

func NewCOS(config *Config) *Adapter {
	a := &Adapter{Config: config}
	a.lock = &sync.Mutex{}
	return a
}

func (a *Adapter) Client() (*cos.Client, error) {
	bucketURL, err := url.Parse(a.Config.BucketURL)
	if err != nil {
		return nil, err
	}
	return cos.NewClient(&cos.BaseURL{BucketURL: bucketURL}, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  a.Config.SecretID,
			SecretKey: a.Config.SecretKey,
		},
	}), nil
}
func (a *Adapter) CopyObject(srcObjectKey, destObjectKey string, isDelete bool) (bool, error) {
	client, err := a.Client()
	if err != nil {
		return false, err
	}
	sourceURL := client.Object.GetObjectURL(srcObjectKey)
	_, resp, err := client.Object.Copy(context.Background(), destObjectKey, sourceURL.Host+sourceURL.Path, nil)
	if err != nil || resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("COS copyObject code=%v ,err=%v", resp.StatusCode, err)
	}
	if isDelete {
		defer func() {
			_, _ = client.Object.Delete(context.Background(), srcObjectKey, nil)
		}()
	}
	return true, nil
}
func (a Adapter) Head(path string) (*cos.Response, error) {
	client, err := a.Client()
	if err != nil {
		return nil, err
	}
	resp, err := client.Object.Head(context.Background(), path, nil)
	if err == nil && resp.StatusCode == http.StatusOK {
		return resp, err
	}
	return nil, fmt.Errorf("COS Head code=%v ,err=%v", resp.StatusCode, err)
}
func (a *Adapter) Exist(path string) (bool, error) {
	client, err := a.Client()
	if err != nil {
		return false, err
	}
	return client.Object.IsExist(context.Background(), path)
}

func (a *Adapter) WriteReader(path string, reader io.Reader) error {
	client, err := a.Client()
	if err != nil {
		return err
	}
	_, err = client.Object.Put(context.Background(), path, reader, nil)
	return err
}

func (a *Adapter) Write(path string, contents []byte) error {
	return a.WriteReader(path, strings.NewReader(string(contents)))
}

func (a *Adapter) WriteStream(path, resource string) error {
	client, err := a.Client()
	if err != nil {
		return err
	}
	_, err = client.Object.PutFromFile(context.Background(), path, resource, nil)
	return err
}

func (a *Adapter) Read(path string) ([]byte, error) {
	client, err := a.Client()
	if err != nil {
		return nil, err
	}
	sourceURL := client.Object.GetObjectURL(path).String()
	resp, err := http.Get(sourceURL)
	if err == nil && resp.StatusCode == http.StatusOK {
		return io.ReadAll(resp.Body)
	}
	return nil, err
}

func (a *Adapter) Delete(path string) (int64, error) {
	client, err := a.Client()
	if err != nil {
		return 0, err
	}
	resp, err := client.Object.Delete(context.Background(), path)
	if err == nil && resp.StatusCode == http.StatusOK {
		return 1, nil
	}
	return 0, err
}

func (a *Adapter) Size(path string) (int64, error) {
	head, err := a.Head(path)
	if err != nil {
		return 0, err
	}
	return head.ContentLength, nil
}

func (a *Adapter) Update(path string, contents []byte) error {
	return a.Write(path, contents)
}

func (a *Adapter) UpdateStream(path, resource string) error {
	return a.WriteStream(path, resource)
}

func (a *Adapter) MimeType(path string) (string, error) {
	head, err := a.Head(path)
	if err != nil {
		return "", err
	}
	return head.Response.Header.Get(v2.HeaderGetContentType), nil
}

func (a *Adapter) Move(source, destination string) (bool, error) {
	return a.CopyObject(source, destination, true)
}

func (a *Adapter) Copy(source, destination string) (bool, error) {
	return a.CopyObject(source, destination, false)
}

func (a *Adapter) DiskName() string {
	return v2.DiskNameCOS
}
