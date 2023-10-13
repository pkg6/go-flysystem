package fslocal

import (
	"fmt"
	"github.com/pkg6/go-flysystem/v2"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"sync"
)

type Adapter struct {
	Config *Config
	lock   *sync.Mutex
}

func New(config v2.IAdapterConfig) v2.IAdapter {
	return config.New()
}

func NewLocal(config *Config) *Adapter {
	a := &Adapter{Config: config}
	a.lock = &sync.Mutex{}
	return a
}

func (f Adapter) DiskName() string {
	return v2.DiskNameLocal
}

func (f *Adapter) Exist(path string) (bool, error) {
	f.lock.Lock()
	defer f.lock.Unlock()
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, err
	}
	return false, err
}
func (f *Adapter) Write(path string, contents []byte) error {
	f.lock.Lock()
	defer f.lock.Unlock()
	dir, err := filepath.Abs(filepath.Dir(path))
	if err != nil {
		return err
	}
	if err = f.ensureDirectory(dir); err != nil {
		return err
	}
	if err = os.WriteFile(path, contents, v2.ModeFilePublic); err != nil {
		return err
	}
	return nil
}

func (f *Adapter) WriteReader(path string, reader io.Reader) error {
	f.lock.Lock()
	defer f.lock.Unlock()
	dir, err := filepath.Abs(filepath.Dir(path))
	if err != nil {
		return err
	}
	if err = f.ensureDirectory(dir); err != nil {
		return err
	}
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = io.Copy(file, reader)
	if err != nil {
		return err
	}
	return nil
}

func (f *Adapter) WriteStream(path, resource string) error {
	contents, err := os.ReadFile(resource)
	if err != nil {
		return err
	}
	return f.Write(path, contents)
}

func (f *Adapter) Update(path string, contents []byte) error {
	f.lock.Lock()
	defer f.lock.Unlock()
	if err := os.WriteFile(path, contents, v2.ModeFilePublic); err != nil {
		return err
	}
	return nil
}

func (f *Adapter) UpdateStream(path, resource string) error {
	contents, err := os.ReadFile(resource)
	if err != nil {
		return err
	}
	return f.Update(path, contents)
}

func (f *Adapter) Read(path string) ([]byte, error) {
	f.lock.Lock()
	defer f.lock.Unlock()
	contents, err := os.ReadFile(path)
	return contents, err
}

func (f *Adapter) Delete(path string) (int64, error) {
	f.lock.Lock()
	defer f.lock.Unlock()
	err := os.Remove(path)
	if err == nil {
		return 0, err
	}
	return 1, nil
}

func (f *Adapter) Size(path string) (int64, error) {
	f.lock.Lock()
	defer f.lock.Unlock()
	info, err := os.Stat(path)
	if err != nil {
		return 0, err
	}
	return info.Size(), err
}

func (f *Adapter) Copy(source, destination string) (bool, error) {
	f.lock.Lock()
	defer f.lock.Unlock()
	info, err := os.Stat(source)
	if err != nil {
		return false, err
	}
	input, err := os.ReadFile(source)
	if err != nil {
		return false, fmt.Errorf("unable to copy file from %s to %s", source, destination)
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

func (f *Adapter) Move(source, destination string) (bool, error) {
	f.lock.Lock()
	defer f.lock.Unlock()
	info, err := os.Stat(source)
	if err != nil {
		return false, err
	}
	input, err := os.ReadFile(source)
	if err != nil {
		return false, fmt.Errorf("unable to copy file from %s to %s", source, destination)
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

// MimeType 可以使用net/http包中提供的DetectContentType函数来获取文件MimeType
func (f *Adapter) MimeType(path string) (string, error) {
	f.lock.Lock()
	defer f.lock.Unlock()
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

func (f *Adapter) CreateDirectory(dirname string) error {
	f.lock.Lock()
	defer f.lock.Unlock()
	if err := f.ensureDirectory(dirname); err != nil {
		return err
	}
	return nil
}

func (f *Adapter) DeleteDirectory(dirname string) (int64, error) {
	f.lock.Lock()
	defer f.lock.Unlock()
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
func (f *Adapter) SetVisibility(path, visibility string) (bool, error) {
	f.lock.Lock()
	defer f.lock.Unlock()
	var err error
	info, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	var permission os.FileMode
	if info.IsDir() {
		permission = v2.FileModes[v2.PathTypeDirectory][visibility]
	} else {
		permission = v2.FileModes[v2.PathTypeFile][visibility]
	}
	err = os.Chmod(path, permission)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (f *Adapter) Visibility(path string) error {
	f.lock.Lock()
	defer f.lock.Unlock()
	_, err := os.Stat(path)
	return err
}
func (f *Adapter) ensureDirectory(root string) error {
	var err error
	if _, err = os.Stat(root); os.IsNotExist(err) {
		if err = os.MkdirAll(root, v2.ModeDirPublic); err != nil {
			return fmt.Errorf("impossible to create directory %s err=%s", root, err.Error())
		}
	}
	return err
}
