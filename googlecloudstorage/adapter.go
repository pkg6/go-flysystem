package googlecloudstorage

import (
	"bytes"
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"github.com/pkg6/go-flysystem"
	"google.golang.org/api/option"
	"io"
	"os"
	"sync"
	"time"
)

var (
	DefaultWithTimeout = time.Second * 50
)

type Config struct {
	Bucket      string
	WithTimeout time.Duration
	Option      []option.ClientOption
}

type Adapter struct {
	Config *Config
	ctx    context.Context
	lock   *sync.Mutex
}

func (a *Adapter) DiskName() string {
	return flysystem.DiskNameGoogleCloudCloudStorage
}

func (a *Adapter) Clone() flysystem.IAdapter {
	if a.Config.WithTimeout == 0 {
		a.Config.WithTimeout = DefaultWithTimeout
	}
	return a
}

func (a *Adapter) StorageClient() (*storage.Client, error) {
	client, err := storage.NewClient(a.ctx, a.Config.Option...)
	if err != nil {
		return nil, err
	}
	defer client.Close()
	return client, err
}

func (a *Adapter) StorageObject(object string) (*storage.ObjectHandle, error) {
	client, err := a.StorageClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()
	return client.Bucket(a.Config.Bucket).Object(object), nil
}

func (a *Adapter) StorageObjectTimeout(object string) (*storage.ObjectHandle, context.Context, error) {
	objectHandle, err := a.StorageObject(object)
	if err != nil {
		return nil, a.ctx, err
	}
	ctx, cancel := context.WithTimeout(a.ctx, a.Config.WithTimeout)
	defer cancel()
	return objectHandle, ctx, nil
}

func (a *Adapter) Exists(path string) (bool, error) {
	size, err := a.Size(path)
	if size > 0 {
		return true, nil
	}
	return false, err
}

func (a *Adapter) WriteReader(path string, reader io.Reader) (string, error) {
	storageObject, ctx, err := a.StorageObjectTimeout(path)
	if err != nil {
		return "", fmt.Errorf("storage.storageObject: %w", err)
	}
	// Upload an object with storage.Writer.
	wc := storageObject.NewWriter(ctx)
	wc.ChunkSize = 0 // note retries are not supported for chunk size 0.
	if _, err = io.Copy(wc, reader); err != nil {
		return "", fmt.Errorf("io.Copy: %w", err)
	}
	// Data can continue to be added to the file until the writer is closed.
	if err := wc.Close(); err != nil {
		return "", fmt.Errorf("Writer.Close: %w", err)
	}
	return path, nil
}

func (a *Adapter) Write(path string, contents []byte) (string, error) {
	return a.WriteReader(path, bytes.NewBuffer(contents))
}

func (a *Adapter) WriteStream(path, resource string) (string, error) {
	f, err := os.Create(resource)
	if err != nil {
		return "", fmt.Errorf("os.Create: %w", err)
	}
	return a.WriteReader(path, f)
}

func (a *Adapter) Read(path string) ([]byte, error) {
	object, ctx, err := a.StorageObjectTimeout(path)
	if err != nil {
		return nil, err
	}
	rc, err := object.NewReader(ctx)
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

func (a *Adapter) Update(path string, contents []byte) (string, error) {
	return a.Write(path, contents)
}

func (a *Adapter) UpdateStream(path, resource string) (string, error) {
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

func (a *Adapter) GetMetadata(object string) (*storage.ObjectAttrs, error) {
	obj, ctx, err := a.StorageObjectTimeout(object)
	if err != nil {
		return nil, err
	}
	return obj.Attrs(ctx)
}
func (a *Adapter) CopyObject(source, destination string, deleteSource bool) (bool, error) {
	src, ctx, err := a.StorageObjectTimeout(source)
	if err != nil {
		return false, err
	}
	dst, _, err := a.StorageObjectTimeout(destination)
	if err != nil {
		return false, err
	}
	dst = dst.If(storage.Conditions{DoesNotExist: true})
	if _, err := dst.CopierFrom(src).Run(ctx); err != nil {
		return false, fmt.Errorf("Object(%q).CopierFrom(%q).Run: %w", source, destination, err)
	}
	if deleteSource {
		if err := src.Delete(ctx); err != nil {
			return false, fmt.Errorf("Object(%q).Delete: %w", source, err)
		}
	}
	return true, nil
}
