package local

import (
	"fmt"
	"github.com/zzqqw/gfs/localfs"
	"io"
	"net/url"
	"os"
	"sync"

	"github.com/pkg6/go-flysystem"
	"github.com/zzqqw/gfs"
)

type Local struct {
	gfs.AbstractAdapter
	root string
	CDN  string
	lock *sync.Mutex
}

func New(root, CDN string) flysystem.IAdapter {
	f := &Local{root: root, CDN: CDN}
	err := f.ensureDirectory(f.root)
	if err != nil {
		panic(err)
	}
	f.lock = &sync.Mutex{}
	f.SetPathPrefix(f.root)
	return f
}

func (f Local) DiskName() string {
	return flysystem.DiskNameLocal
}

func (f *Local) GFSAdapter() gfs.IAdapter {
	return localfs.New(&localfs.Config{CDN: f.CDN})
}

func (f *Local) URL(path string) (*url.URL, error) {
	path = f.ApplyPathPrefix(path)
	return f.GFSAdapter().URL(path)
}
func (f *Local) Exists(path string) (bool, error) {
	path = f.ApplyPathPrefix(path)
	return f.GFSAdapter().Exist(path)
}

func (f *Local) WriteReader(path string, reader io.Reader) (string, error) {
	path = f.ApplyPathPrefix(path)
	err := f.GFSAdapter().WriteReader(path, reader)
	return path, err
}

func (f *Local) Write(path string, contents []byte) (string, error) {
	path = f.ApplyPathPrefix(path)
	err := f.GFSAdapter().Write(path, contents)
	return path, err
}

func (f *Local) WriteStream(path, resource string) (string, error) {
	path = f.ApplyPathPrefix(path)
	err := f.GFSAdapter().WriteStream(path, resource)
	return path, err
}

func (f *Local) Read(path string) ([]byte, error) {
	path = f.ApplyPathPrefix(path)
	return f.GFSAdapter().Read(path)
}

func (f *Local) Delete(path string) (int64, error) {
	path = f.ApplyPathPrefix(path)
	return f.GFSAdapter().Delete(path)
}

func (f *Local) Size(path string) (int64, error) {
	path = f.ApplyPathPrefix(path)
	return f.GFSAdapter().Size(path)
}

func (f *Local) Update(path string, contents []byte) (string, error) {
	path = f.ApplyPathPrefix(path)
	err := f.GFSAdapter().Update(path, contents)
	return path, err
}

func (f *Local) UpdateStream(path, resource string) (string, error) {
	path = f.ApplyPathPrefix(path)
	err := f.GFSAdapter().UpdateStream(path, resource)
	return path, err
}

func (f *Local) MimeType(path string) (string, error) {
	path = f.ApplyPathPrefix(path)
	return f.GFSAdapter().MimeType(path)
}

func (f *Local) Move(source, destination string) (bool, error) {
	source = f.ApplyPathPrefix(source)
	destination = f.ApplyPathPrefix(destination)
	return f.GFSAdapter().Move(source, destination)
}

func (f *Local) Copy(source, destination string) (bool, error) {
	source = f.ApplyPathPrefix(source)
	destination = f.ApplyPathPrefix(destination)
	return f.GFSAdapter().Copy(source, destination)
}
func (f *Local) ensureDirectory(root string) error {
	var err error
	if _, err = os.Stat(root); os.IsNotExist(err) {
		if err = os.MkdirAll(root, gfs.ModeDirPublic); err != nil {
			return fmt.Errorf("impossible to create directory %s err=%s", root, err.Error())
		}
	}
	return err
}
