package fscos

import (
	"github.com/pkg6/go-flysystem"
	"github.com/pkg6/gfs"
	"io"
	"net/url"
	"reflect"
	"sync"
	"testing"
)

func TestFSCos_Copy(t *testing.T) {
	type fields struct {
		AbstractAdapter gfs.AbstractAdapter
		Config          *Config
		lock            *sync.Mutex
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
			a := &FSCos{
				AbstractAdapter: tt.fields.AbstractAdapter,
				Config:          tt.fields.Config,
				lock:            tt.fields.lock,
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

func TestFSCos_Delete(t *testing.T) {
	type fields struct {
		AbstractAdapter gfs.AbstractAdapter
		Config          *Config
		lock            *sync.Mutex
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
			a := &FSCos{
				AbstractAdapter: tt.fields.AbstractAdapter,
				Config:          tt.fields.Config,
				lock:            tt.fields.lock,
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

func TestFSCos_DiskName(t *testing.T) {
	type fields struct {
		AbstractAdapter gfs.AbstractAdapter
		Config          *Config
		lock            *sync.Mutex
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
			a := &FSCos{
				AbstractAdapter: tt.fields.AbstractAdapter,
				Config:          tt.fields.Config,
				lock:            tt.fields.lock,
			}
			if got := a.DiskName(); got != tt.want {
				t.Errorf("DiskName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFSCos_Exists(t *testing.T) {
	type fields struct {
		AbstractAdapter gfs.AbstractAdapter
		Config          *Config
		lock            *sync.Mutex
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
			a := &FSCos{
				AbstractAdapter: tt.fields.AbstractAdapter,
				Config:          tt.fields.Config,
				lock:            tt.fields.lock,
			}
			got, err := a.Exists(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("Exists() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Exists() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFSCos_GFSAdapter(t *testing.T) {
	type fields struct {
		AbstractAdapter gfs.AbstractAdapter
		Config          *Config
		lock            *sync.Mutex
	}
	tests := []struct {
		name   string
		fields fields
		want   gfs.IAdapter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &FSCos{
				AbstractAdapter: tt.fields.AbstractAdapter,
				Config:          tt.fields.Config,
				lock:            tt.fields.lock,
			}
			if got := a.GFSAdapter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GFSAdapter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFSCos_MimeType(t *testing.T) {
	type fields struct {
		AbstractAdapter gfs.AbstractAdapter
		Config          *Config
		lock            *sync.Mutex
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
			a := &FSCos{
				AbstractAdapter: tt.fields.AbstractAdapter,
				Config:          tt.fields.Config,
				lock:            tt.fields.lock,
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

func TestFSCos_Move(t *testing.T) {
	type fields struct {
		AbstractAdapter gfs.AbstractAdapter
		Config          *Config
		lock            *sync.Mutex
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
			a := &FSCos{
				AbstractAdapter: tt.fields.AbstractAdapter,
				Config:          tt.fields.Config,
				lock:            tt.fields.lock,
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

func TestFSCos_Read(t *testing.T) {
	type fields struct {
		AbstractAdapter gfs.AbstractAdapter
		Config          *Config
		lock            *sync.Mutex
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
			a := &FSCos{
				AbstractAdapter: tt.fields.AbstractAdapter,
				Config:          tt.fields.Config,
				lock:            tt.fields.lock,
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

func TestFSCos_Size(t *testing.T) {
	type fields struct {
		AbstractAdapter gfs.AbstractAdapter
		Config          *Config
		lock            *sync.Mutex
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
			a := &FSCos{
				AbstractAdapter: tt.fields.AbstractAdapter,
				Config:          tt.fields.Config,
				lock:            tt.fields.lock,
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

func TestFSCos_URL(t *testing.T) {
	type fields struct {
		AbstractAdapter gfs.AbstractAdapter
		Config          *Config
		lock            *sync.Mutex
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
			a := &FSCos{
				AbstractAdapter: tt.fields.AbstractAdapter,
				Config:          tt.fields.Config,
				lock:            tt.fields.lock,
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

func TestFSCos_Update(t *testing.T) {
	type fields struct {
		AbstractAdapter gfs.AbstractAdapter
		Config          *Config
		lock            *sync.Mutex
	}
	type args struct {
		path     string
		contents []byte
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
			a := &FSCos{
				AbstractAdapter: tt.fields.AbstractAdapter,
				Config:          tt.fields.Config,
				lock:            tt.fields.lock,
			}
			got, err := a.Update(tt.args.path, tt.args.contents)
			if (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Update() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFSCos_UpdateStream(t *testing.T) {
	type fields struct {
		AbstractAdapter gfs.AbstractAdapter
		Config          *Config
		lock            *sync.Mutex
	}
	type args struct {
		path     string
		resource string
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
			a := &FSCos{
				AbstractAdapter: tt.fields.AbstractAdapter,
				Config:          tt.fields.Config,
				lock:            tt.fields.lock,
			}
			got, err := a.UpdateStream(tt.args.path, tt.args.resource)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateStream() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UpdateStream() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFSCos_Write(t *testing.T) {
	type fields struct {
		AbstractAdapter gfs.AbstractAdapter
		Config          *Config
		lock            *sync.Mutex
	}
	type args struct {
		path     string
		contents []byte
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
			a := &FSCos{
				AbstractAdapter: tt.fields.AbstractAdapter,
				Config:          tt.fields.Config,
				lock:            tt.fields.lock,
			}
			got, err := a.Write(tt.args.path, tt.args.contents)
			if (err != nil) != tt.wantErr {
				t.Errorf("Write() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Write() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFSCos_WriteReader(t *testing.T) {
	type fields struct {
		AbstractAdapter gfs.AbstractAdapter
		Config          *Config
		lock            *sync.Mutex
	}
	type args struct {
		path   string
		reader io.Reader
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
			a := &FSCos{
				AbstractAdapter: tt.fields.AbstractAdapter,
				Config:          tt.fields.Config,
				lock:            tt.fields.lock,
			}
			got, err := a.WriteReader(tt.args.path, tt.args.reader)
			if (err != nil) != tt.wantErr {
				t.Errorf("WriteReader() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("WriteReader() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFSCos_WriteStream(t *testing.T) {
	type fields struct {
		AbstractAdapter gfs.AbstractAdapter
		Config          *Config
		lock            *sync.Mutex
	}
	type args struct {
		path     string
		resource string
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
			a := &FSCos{
				AbstractAdapter: tt.fields.AbstractAdapter,
				Config:          tt.fields.Config,
				lock:            tt.fields.lock,
			}
			got, err := a.WriteStream(tt.args.path, tt.args.resource)
			if (err != nil) != tt.wantErr {
				t.Errorf("WriteStream() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("WriteStream() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		config *Config
	}
	tests := []struct {
		name string
		args args
		want flysystem.IAdapter
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
