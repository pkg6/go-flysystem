package fslocal

import (
	"fmt"
	v2 "github.com/pkg6/go-flysystem/v2"
	"io"
	"net/http"
	"net/url"
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

func (a Adapter) DiskName() string {
	return v2.DiskNameLocal
}
func (a *Adapter) URL(path string) (*url.URL, error) {
	return a.Config.URL(path)
}
func (a *Adapter) Exist(path string) (bool, error) {
	a.lock.Lock()
	defer a.lock.Unlock()
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, err
	}
	return false, err
}
func (a *Adapter) Write(path string, contents []byte) error {
	a.lock.Lock()
	defer a.lock.Unlock()
	dir, err := filepath.Abs(filepath.Dir(path))
	if err != nil {
		return err
	}
	if err = a.ensureDirectory(dir); err != nil {
		return err
	}
	if err = os.WriteFile(path, contents, v2.ModeFilePublic); err != nil {
		return err
	}
	return nil
}

func (a *Adapter) WriteReader(path string, reader io.Reader) error {
	a.lock.Lock()
	defer a.lock.Unlock()
	dir, err := filepath.Abs(filepath.Dir(path))
	if err != nil {
		return err
	}
	if err = a.ensureDirectory(dir); err != nil {
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

func (a *Adapter) WriteStream(path, resource string) error {
	contents, err := os.ReadFile(resource)
	if err != nil {
		return err
	}
	return a.Write(path, contents)
}

func (a *Adapter) Update(path string, contents []byte) error {
	a.lock.Lock()
	defer a.lock.Unlock()
	if err := os.WriteFile(path, contents, v2.ModeFilePublic); err != nil {
		return err
	}
	return nil
}

func (a *Adapter) UpdateStream(path, resource string) error {
	contents, err := os.ReadFile(resource)
	if err != nil {
		return err
	}
	return a.Update(path, contents)
}

func (a *Adapter) Read(path string) ([]byte, error) {
	a.lock.Lock()
	defer a.lock.Unlock()
	contents, err := os.ReadFile(path)
	return contents, err
}

func (a *Adapter) Delete(path string) (int64, error) {
	a.lock.Lock()
	defer a.lock.Unlock()
	err := os.Remove(path)
	if err == nil {
		return 0, err
	}
	return 1, nil
}

func (a *Adapter) Size(path string) (int64, error) {
	a.lock.Lock()
	defer a.lock.Unlock()
	info, err := os.Stat(path)
	if err != nil {
		return 0, err
	}
	return info.Size(), err
}

func (a *Adapter) Copy(source, destination string) (bool, error) {
	a.lock.Lock()
	defer a.lock.Unlock()
	info, err := os.Stat(source)
	if err != nil {
		return false, err
	}
	input, err := os.ReadFile(source)
	if err != nil {
		return false, fmt.Errorf("unable to copy file from %s to %s", source, destination)
	}
	if err := a.ensureDirectory(path.Dir(destination)); err != nil {
		return false, err
	}
	err = os.WriteFile(destination, input, info.Mode())
	if err != nil {
		return false, err
	}
	return true, nil
}

func (a *Adapter) Move(source, destination string) (bool, error) {
	a.lock.Lock()
	defer a.lock.Unlock()
	info, err := os.Stat(source)
	if err != nil {
		return false, err
	}
	input, err := os.ReadFile(source)
	if err != nil {
		return false, fmt.Errorf("unable to copy file from %s to %s", source, destination)
	}
	if err := a.ensureDirectory(path.Dir(destination)); err != nil {
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
func (a *Adapter) MimeType(path string) (string, error) {
	a.lock.Lock()
	defer a.lock.Unlock()
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

func (a *Adapter) CreateDirectory(dirname string) error {
	a.lock.Lock()
	defer a.lock.Unlock()
	if err := a.ensureDirectory(dirname); err != nil {
		return err
	}
	return nil
}

func (a *Adapter) DeleteDirectory(dirname string) (int64, error) {
	a.lock.Lock()
	defer a.lock.Unlock()
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
func (a *Adapter) SetVisibility(path, visibility string) (bool, error) {
	a.lock.Lock()
	defer a.lock.Unlock()
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

func (a *Adapter) Visibility(path string) error {
	a.lock.Lock()
	defer a.lock.Unlock()
	_, err := os.Stat(path)
	return err
}
func (a *Adapter) ensureDirectory(root string) error {
	var err error
	if _, err = os.Stat(root); os.IsNotExist(err) {
		if err = os.MkdirAll(root, v2.ModeDirPublic); err != nil {
			return fmt.Errorf("impossible to create directory %s err=%s", root, err.Error())
		}
	}
	return err
}
