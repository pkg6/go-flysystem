package local

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"sync"

	"github.com/pkg6/go-flysystem"
	"github.com/zzqqw/gfs"
)

type Local struct {
	gfs.AbstractAdapter
	root string
	lock *sync.Mutex
}

func New(root string) flysystem.IAdapter {
	l := Local{root: root}
	return l.Clone()
}

func (f Local) DiskName() string {
	return flysystem.DiskNameLocal
}

func (f Local) Clone() flysystem.IAdapter {
	err := f.ensureDirectory(f.root)
	if err != nil {
		panic(err)
	}
	f.lock = &sync.Mutex{}
	f.SetPathPrefix(f.root)
	return &f
}

func (f *Local) URL(path string) (*url.URL, error) {
	path = f.ApplyPathPrefix(path)
	return nil, fmt.Errorf("url nill")
}
func (f *Local) Exists(path string) (bool, error) {
	f.lock.Lock()
	defer f.lock.Unlock()
	path = f.ApplyPathPrefix(path)
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, err
	}
	return false, err
}
func (f *Local) Write(path string, contents []byte) (string, error) {
	f.lock.Lock()
	defer f.lock.Unlock()
	path = f.ApplyPathPrefix(path)
	dir, err := filepath.Abs(filepath.Dir(path))
	if err != nil {
		return "", err
	}
	if err = f.ensureDirectory(dir); err != nil {
		return "", err
	}
	if err = os.WriteFile(path, contents, gfs.ModeFilePublic); err != nil {
		return "", err
	}
	return path, nil
}

func (f *Local) WriteReader(path string, reader io.Reader) (string, error) {
	f.lock.Lock()
	defer f.lock.Unlock()
	path = f.ApplyPathPrefix(path)
	dir, err := filepath.Abs(filepath.Dir(path))
	if err != nil {
		return "", err
	}
	if err = f.ensureDirectory(dir); err != nil {
		return "", err
	}
	file, err := os.Create(path)
	if err != nil {
		return "", err
	}
	defer file.Close()
	_, err = io.Copy(file, reader)
	if err != nil {
		return "", err
	}
	return file.Name(), nil
}

func (f *Local) WriteStream(path, resource string) (string, error) {
	contents, err := os.ReadFile(resource)
	if err != nil {
		return "", err
	}
	return f.Write(path, contents)
}

func (f *Local) Update(path string, contents []byte) (string, error) {
	f.lock.Lock()
	defer f.lock.Unlock()
	path = f.ApplyPathPrefix(path)
	if err := os.WriteFile(path, contents, gfs.ModeFilePublic); err != nil {
		return "", err
	}
	return path, nil
}

func (f *Local) UpdateStream(path, resource string) (string, error) {
	contents, err := os.ReadFile(resource)
	if err != nil {
		return "", err
	}
	return f.Update(path, contents)
}

func (f *Local) Read(path string) ([]byte, error) {
	f.lock.Lock()
	defer f.lock.Unlock()
	path = f.ApplyPathPrefix(path)
	contents, err := os.ReadFile(path)
	return contents, err
}

func (f *Local) Delete(path string) (int64, error) {
	f.lock.Lock()
	defer f.lock.Unlock()
	path = f.ApplyPathPrefix(path)
	err := os.Remove(path)
	if err == nil {
		return 0, err
	}
	return 1, nil
}

func (f *Local) CreateDirectory(dirname string) error {
	f.lock.Lock()
	defer f.lock.Unlock()
	dirname = f.ApplyPathPrefix(dirname)
	if err := f.ensureDirectory(dirname); err != nil {
		return err
	}
	return nil
}

func (f *Local) DeleteDirectory(dirname string) (int64, error) {
	f.lock.Lock()
	defer f.lock.Unlock()
	dirname = f.ApplyPathPrefix(dirname)
	var err error
	var count int64
	if err = filepath.Walk(dirname, func(path string, info os.FileInfo, err error) error {
		count++
		return nil
	}); err != nil {
		return 0, err
	}
	if err = os.RemoveAll(dirname); err != nil {
		return 0, err
	}
	return count, nil
}

func (f *Local) Size(path string) (int64, error) {
	f.lock.Lock()
	defer f.lock.Unlock()
	path = f.ApplyPathPrefix(path)
	info, err := os.Stat(path)
	if err != nil {
		return 0, err
	}
	return info.Size(), err
}

func (f *Local) Copy(source, destination string) (bool, error) {
	f.lock.Lock()
	defer f.lock.Unlock()
	source = f.ApplyPathPrefix(source)
	destination = f.ApplyPathPrefix(destination)
	info, err := os.Stat(source)
	if err != nil {
		return false, err
	}
	input, err := os.ReadFile(source)
	if err != nil {
		return false, errors.New("Unable to copy file from " + source + " to " + destination)
	}
	if err := f.ensureDirectory(path.Dir(destination)); err != nil {
		return false, err
	}
	err = os.WriteFile(destination, input, info.Mode())
	if err != nil {
		return false, err
	}
	return true, nil
}

func (f *Local) Move(source, destination string) (bool, error) {
	f.lock.Lock()
	defer f.lock.Unlock()
	source = f.ApplyPathPrefix(source)
	destination = f.ApplyPathPrefix(destination)
	info, err := os.Stat(source)
	if err != nil {
		return false, err
	}
	input, err := os.ReadFile(source)
	if err != nil {
		return false, errors.New("Unable to copy file from " + source + " to " + destination)
	}
	if err := f.ensureDirectory(path.Dir(destination)); err != nil {
		return false, err
	}
	err = os.WriteFile(destination, input, info.Mode())
	if err != nil {
		return false, err
	}
	err = os.Remove(source)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (f *Local) SetVisibility(path, visibility string) (bool, error) {
	f.lock.Lock()
	defer f.lock.Unlock()
	var err error
	path = f.ApplyPathPrefix(path)
	info, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	var permission os.FileMode
	if info.IsDir() {
		permission = gfs.FileModes[gfs.PathTypeDirectory][visibility]
	} else {
		permission = gfs.FileModes[gfs.PathTypeFile][visibility]
	}
	err = os.Chmod(path, permission)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (f *Local) Visibility(path string) error {
	f.lock.Lock()
	defer f.lock.Unlock()
	path = f.ApplyPathPrefix(path)
	_, err := os.Stat(path)
	return err
}

// MimeType 可以使用net/http包中提供的DetectContentType函数来获取文件MimeType
func (f *Local) MimeType(path string) (string, error) {
	f.lock.Lock()
	defer f.lock.Unlock()
	path = f.ApplyPathPrefix(path)
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()
	//获取文件MimeType
	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil {
		return "", err
	}
	return http.DetectContentType(buffer), nil
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
