package fslocal

import (
	"github.com/pkg6/go-flysystem/gfs"
	"io"
	"net/url"
	"reflect"
	"sync"
	"testing"
)

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
			f := &Adapter{
				Config: tt.fields.Config,
				lock:   tt.fields.lock,
			}
			got, err := f.Copy(tt.args.source, tt.args.destination)
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

func TestAdapter_CreateDirectory(t *testing.T) {
	type fields struct {
		Config *Config
		lock   *sync.Mutex
	}
	type args struct {
		dirname string
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
			f := &Adapter{
				Config: tt.fields.Config,
				lock:   tt.fields.lock,
			}
			if err := f.CreateDirectory(tt.args.dirname); (err != nil) != tt.wantErr {
				t.Errorf("CreateDirectory() error = %v, wantErr %v", err, tt.wantErr)
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
			f := &Adapter{
				Config: tt.fields.Config,
				lock:   tt.fields.lock,
			}
			got, err := f.Delete(tt.args.path)
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

func TestAdapter_DeleteDirectory(t *testing.T) {
	type fields struct {
		Config *Config
		lock   *sync.Mutex
	}
	type args struct {
		dirname string
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
			f := &Adapter{
				Config: tt.fields.Config,
				lock:   tt.fields.lock,
			}
			got, err := f.DeleteDirectory(tt.args.dirname)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteDirectory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DeleteDirectory() got = %v, want %v", got, tt.want)
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
			f := Adapter{
				Config: tt.fields.Config,
				lock:   tt.fields.lock,
			}
			if got := f.DiskName(); got != tt.want {
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
			f := &Adapter{
				Config: tt.fields.Config,
				lock:   tt.fields.lock,
			}
			got, err := f.Exist(tt.args.path)
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
			f := &Adapter{
				Config: tt.fields.Config,
				lock:   tt.fields.lock,
			}
			got, err := f.MimeType(tt.args.path)
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
			f := &Adapter{
				Config: tt.fields.Config,
				lock:   tt.fields.lock,
			}
			got, err := f.Move(tt.args.source, tt.args.destination)
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
			f := &Adapter{
				Config: tt.fields.Config,
				lock:   tt.fields.lock,
			}
			got, err := f.Read(tt.args.path)
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

func TestAdapter_SetVisibility(t *testing.T) {
	type fields struct {
		Config *Config
		lock   *sync.Mutex
	}
	type args struct {
		path       string
		visibility string
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
			f := &Adapter{
				Config: tt.fields.Config,
				lock:   tt.fields.lock,
			}
			got, err := f.SetVisibility(tt.args.path, tt.args.visibility)
			if (err != nil) != tt.wantErr {
				t.Errorf("SetVisibility() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SetVisibility() got = %v, want %v", got, tt.want)
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
			f := &Adapter{
				Config: tt.fields.Config,
				lock:   tt.fields.lock,
			}
			got, err := f.Size(tt.args.path)
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
			f := &Adapter{
				Config: tt.fields.Config,
				lock:   tt.fields.lock,
			}
			if err := f.Update(tt.args.path, tt.args.contents); (err != nil) != tt.wantErr {
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
			f := &Adapter{
				Config: tt.fields.Config,
				lock:   tt.fields.lock,
			}
			if err := f.UpdateStream(tt.args.path, tt.args.resource); (err != nil) != tt.wantErr {
				t.Errorf("UpdateStream() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAdapter_Visibility(t *testing.T) {
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
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Adapter{
				Config: tt.fields.Config,
				lock:   tt.fields.lock,
			}
			if err := f.Visibility(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("Visibility() error = %v, wantErr %v", err, tt.wantErr)
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
			f := &Adapter{
				Config: tt.fields.Config,
				lock:   tt.fields.lock,
			}
			if err := f.Write(tt.args.path, tt.args.contents); (err != nil) != tt.wantErr {
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
			f := &Adapter{
				Config: tt.fields.Config,
				lock:   tt.fields.lock,
			}
			if err := f.WriteReader(tt.args.path, tt.args.reader); (err != nil) != tt.wantErr {
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
			f := &Adapter{
				Config: tt.fields.Config,
				lock:   tt.fields.lock,
			}
			if err := f.WriteStream(tt.args.path, tt.args.resource); (err != nil) != tt.wantErr {
				t.Errorf("WriteStream() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAdapter_ensureDirectory(t *testing.T) {
	type fields struct {
		Config *Config
		lock   *sync.Mutex
	}
	type args struct {
		root string
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
			f := &Adapter{
				Config: tt.fields.Config,
				lock:   tt.fields.lock,
			}
			if err := f.ensureDirectory(tt.args.root); (err != nil) != tt.wantErr {
				t.Errorf("ensureDirectory() error = %v, wantErr %v", err, tt.wantErr)
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

func TestNewLocal(t *testing.T) {
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
			if got := NewLocal(tt.args.config); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLocal() = %v, want %v", got, tt.want)
			}
		})
	}
}
