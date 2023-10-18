package fskodo

import (
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"github.com/pkg6/go-flysystem/gfs"
	"io"
	"net/url"
	"reflect"
	"sync"
	"testing"
)

func TestAdapter_BucketManager(t *testing.T) {
	type fields struct {
		Config *Config
		lock   *sync.Mutex
	}
	tests := []struct {
		name   string
		fields fields
		want   *storage.BucketManager
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Adapter{
				Config: tt.fields.Config,
				lock:   tt.fields.lock,
			}
			if got := a.BucketManager(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BucketManager() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdapter_BucketManagerBatch(t *testing.T) {
	type fields struct {
		Config *Config
		lock   *sync.Mutex
	}
	type args struct {
		operations []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Adapter{
				Config: tt.fields.Config,
				lock:   tt.fields.lock,
			}
			if err := a.BucketManagerBatch(tt.args.operations); (err != nil) != tt.wantErr {
				t.Errorf("BucketManagerBatch() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAdapter_Copy(t *testing.T) {
	type fields struct {
		Config *Config
		lock   *sync.Mutex
	}
	type args struct {
		source      string
		destination string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Adapter{
				Config: tt.fields.Config,
				lock:   tt.fields.lock,
			}
			got, err := a.Copy(tt.args.source, tt.args.destination)
			if (err != nil) != tt.wantErr {
				t.Errorf("Copy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Copy() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdapter_Delete(t *testing.T) {
	type fields struct {
		Config *Config
		lock   *sync.Mutex
	}
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Adapter{
				Config: tt.fields.Config,
				lock:   tt.fields.lock,
			}
			got, err := a.Delete(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Delete() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdapter_DiskName(t *testing.T) {
	type fields struct {
		Config *Config
		lock   *sync.Mutex
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Adapter{
				Config: tt.fields.Config,
				lock:   tt.fields.lock,
			}
			if got := a.DiskName(); got != tt.want {
				t.Errorf("DiskName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdapter_Exist(t *testing.T) {
	type fields struct {
		Config *Config
		lock   *sync.Mutex
	}
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Adapter{
				Config: tt.fields.Config,
				lock:   tt.fields.lock,
			}
			got, err := a.Exist(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("Exist() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Exist() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdapter_Mac(t *testing.T) {
	type fields struct {
		Config *Config
		lock   *sync.Mutex
	}
	tests := []struct {
		name   string
		fields fields
		want   *qbox.Mac
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Adapter{
				Config: tt.fields.Config,
				lock:   tt.fields.lock,
			}
			if got := a.Mac(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Mac() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdapter_MimeType(t *testing.T) {
	type fields struct {
		Config *Config
		lock   *sync.Mutex
	}
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Adapter{
				Config: tt.fields.Config,
				lock:   tt.fields.lock,
			}
			got, err := a.MimeType(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("MimeType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MimeType() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdapter_Move(t *testing.T) {
	type fields struct {
		Config *Config
		lock   *sync.Mutex
	}
	type args struct {
		source      string
		destination string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Adapter{
				Config: tt.fields.Config,
				lock:   tt.fields.lock,
			}
			got, err := a.Move(tt.args.source, tt.args.destination)
			if (err != nil) != tt.wantErr {
				t.Errorf("Move() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Move() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdapter_Read(t *testing.T) {
	type fields struct {
		Config *Config
		lock   *sync.Mutex
	}
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Adapter{
				Config: tt.fields.Config,
				lock:   tt.fields.lock,
			}
			got, err := a.Read(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Read() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdapter_Size(t *testing.T) {
	type fields struct {
		Config *Config
		lock   *sync.Mutex
	}
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Adapter{
				Config: tt.fields.Config,
				lock:   tt.fields.lock,
			}
			got, err := a.Size(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("Size() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Size() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdapter_Stat(t *testing.T) {
	type fields struct {
		Config *Config
		lock   *sync.Mutex
	}
	type args struct {
		path string
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantInfo storage.FileInfo
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Adapter{
				Config: tt.fields.Config,
				lock:   tt.fields.lock,
			}
			gotInfo, err := a.Stat(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("Stat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotInfo, tt.wantInfo) {
				t.Errorf("Stat() gotInfo = %v, want %v", gotInfo, tt.wantInfo)
			}
		})
	}
}

func TestAdapter_StorageConfig(t *testing.T) {
	type fields struct {
		Config *Config
		lock   *sync.Mutex
	}
	tests := []struct {
		name   string
		fields fields
		want   *storage.Config
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Adapter{
				Config: tt.fields.Config,
				lock:   tt.fields.lock,
			}
			if got := a.StorageConfig(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StorageConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdapter_URL(t *testing.T) {
	type fields struct {
		Config *Config
		lock   *sync.Mutex
	}
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *url.URL
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Adapter{
				Config: tt.fields.Config,
				lock:   tt.fields.lock,
			}
			got, err := a.URL(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("URL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("URL() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdapter_Update(t *testing.T) {
	type fields struct {
		Config *Config
		lock   *sync.Mutex
	}
	type args struct {
		path     string
		contents []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Adapter{
				Config: tt.fields.Config,
				lock:   tt.fields.lock,
			}
			if err := a.Update(tt.args.path, tt.args.contents); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAdapter_UpdateStream(t *testing.T) {
	type fields struct {
		Config *Config
		lock   *sync.Mutex
	}
	type args struct {
		path     string
		resource string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Adapter{
				Config: tt.fields.Config,
				lock:   tt.fields.lock,
			}
			if err := a.UpdateStream(tt.args.path, tt.args.resource); (err != nil) != tt.wantErr {
				t.Errorf("UpdateStream() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAdapter_UploadToken(t *testing.T) {
	type fields struct {
		Config *Config
		lock   *sync.Mutex
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Adapter{
				Config: tt.fields.Config,
				lock:   tt.fields.lock,
			}
			if got := a.UploadToken(); got != tt.want {
				t.Errorf("UploadToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdapter_Write(t *testing.T) {
	type fields struct {
		Config *Config
		lock   *sync.Mutex
	}
	type args struct {
		path     string
		contents []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Adapter{
				Config: tt.fields.Config,
				lock:   tt.fields.lock,
			}
			if err := a.Write(tt.args.path, tt.args.contents); (err != nil) != tt.wantErr {
				t.Errorf("Write() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAdapter_WriteReader(t *testing.T) {
	type fields struct {
		Config *Config
		lock   *sync.Mutex
	}
	type args struct {
		path   string
		reader io.Reader
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Adapter{
				Config: tt.fields.Config,
				lock:   tt.fields.lock,
			}
			if err := a.WriteReader(tt.args.path, tt.args.reader); (err != nil) != tt.wantErr {
				t.Errorf("WriteReader() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAdapter_WriteStream(t *testing.T) {
	type fields struct {
		Config *Config
		lock   *sync.Mutex
	}
	type args struct {
		path     string
		resource string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Adapter{
				Config: tt.fields.Config,
				lock:   tt.fields.lock,
			}
			if err := a.WriteStream(tt.args.path, tt.args.resource); (err != nil) != tt.wantErr {
				t.Errorf("WriteStream() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		config gfs.IAdapterConfig
	}
	tests := []struct {
		name string
		args args
		want gfs.IAdapter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.config); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewKoDo(t *testing.T) {
	type args struct {
		config *Config
	}
	tests := []struct {
		name string
		args args
		want *Adapter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewKoDo(tt.args.config); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewKoDo() = %v, want %v", got, tt.want)
			}
		})
	}
}
