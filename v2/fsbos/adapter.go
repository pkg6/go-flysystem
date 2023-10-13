package fsbos

import (
	"github.com/baidubce/bce-sdk-go/services/bos"
	"github.com/baidubce/bce-sdk-go/services/bos/api"
	"github.com/pkg6/go-flysystem/v2"
	"io"
	"net/http"
	"sync"
)

const (
	BJEndpoint  = "https://bj.bcebos.com"
	BDEndpoint  = "https://bd.bcebos.com"
	SUEndpoint  = "https://su.bcebos.com"
	GZEndpoint  = "https://gz.bcebos.com"
	CDEndpoint  = "https://cd.bcebos.com"
	HKGEndpoint = "https://hkg.bcebos.com"
	FWHEndpoint = "https://fwh.bcebos.com"
	FSHEndpoint = "https://fsh.bcebos.com"
)

type Adapter struct {
	Config *Config
	lock   *sync.Mutex
}

func New(config v2.IAdapterConfig) v2.IAdapter {
	return config.New()
}

func NewBOS(config *Config) *Adapter {
	a := &Adapter{Config: config}
	if a.Config.Endpoint == "" {
		a.Config.Endpoint = BJEndpoint
	}
	a.lock = &sync.Mutex{}
	return a
}

func (a *Adapter) Client() (*bos.Client, error) {
	return bos.NewClientWithConfig(&bos.BosClientConfiguration{
		Ak:               a.Config.Ak,
		Sk:               a.Config.Sk,
		Endpoint:         a.Config.Endpoint,
		RedirectDisabled: a.Config.RedirectDisabled,
	})
}

func (a *Adapter) ObjectMeta(path string) (*api.GetObjectMetaResult, error) {
	client, err := a.Client()
	if err != nil {
		return nil, err
	}
	return client.GetObjectMeta(a.Config.Bucket, path)
}

func (a *Adapter) Exist(path string) (bool, error) {
	resp, err := a.ObjectMeta(path)
	if err == nil && resp.ContentMD5 != "" {
		return true, nil
	}
	return false, err
}

func (a *Adapter) WriteReader(path string, reader io.Reader) error {
	contents, err := io.ReadAll(reader)
	if err != nil {
		return err
	}
	return a.Write(path, contents)
}

func (a *Adapter) Write(path string, contents []byte) error {
	client, err := a.Client()
	if err != nil {
		return err
	}
	_, err = client.PutObjectFromBytes(a.Config.Bucket, path, contents, nil)
	return err
}

func (a *Adapter) WriteStream(path, resource string) error {
	client, err := a.Client()
	if err != nil {
		return err
	}
	_, err = client.PutObjectFromFile(a.Config.Bucket, path, resource, nil)
	return err
}

func (a *Adapter) Read(path string) ([]byte, error) {
	uri, err := a.Config.URI(path)
	if err != nil {
		return nil, err
	}
	resp, err := http.Get(uri.String())
	if err != nil || resp.StatusCode != http.StatusOK {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

func (a *Adapter) Delete(path string) (int64, error) {
	client, err := a.Client()
	if err != nil {
		return 0, err
	}
	if err = client.DeleteObject(a.Config.Bucket, path); err == nil {
		return 1, nil
	}
	return 0, err
}

func (a *Adapter) Size(path string) (int64, error) {
	meta, err := a.ObjectMeta(path)
	if err != nil {
		return 0, err
	}
	return meta.ContentLength, nil
}

func (a *Adapter) Update(path string, contents []byte) error {
	return a.Write(path, contents)
}

func (a *Adapter) UpdateStream(path, resource string) error {
	return a.WriteStream(path, resource)
}

func (a *Adapter) MimeType(path string) (string, error) {
	meta, err := a.ObjectMeta(path)
	if err != nil {
		return "", err
	}
	return meta.ContentType, nil
}

func (a *Adapter) Move(source, destination string) (bool, error) {
	return a.CopyObject(source, destination, true)
}

func (a *Adapter) Copy(source, destination string) (bool, error) {
	return a.CopyObject(source, destination, false)
}

func (a *Adapter) CopyObject(source, destination string, deleteSource bool) (bool, error) {
	client, err := a.Client()
	if err != nil {
		return false, err
	}
	if _, err = client.BasicCopyObject(a.Config.Bucket, destination, a.Config.Bucket, source); err == nil {
		if deleteSource {
			defer func() {
				_ = client.DeleteObject(a.Config.Bucket, source)
			}()
		}
		return true, nil
	}
	return false, err
}

func (a *Adapter) DiskName() string {
	return v2.DiskNameBOS
}
