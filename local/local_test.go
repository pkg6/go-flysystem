package local

import (
	"github.com/pkg6/go-flysystem"
	"github.com/zzqqw/gfs"
	"io"
	"net/url"
	"reflect"
	"sync"
	"testing"
)

func TestLocal_Copy(t *testing.T) {
	type fields struct {
		AbstractAdapter gfs.AbstractAdapter
		root            string
		CDN             string
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
			f := &Local{
				AbstractAdapter: tt.fields.AbstractAdapter,
				root:            tt.fields.root,
				CDN:             tt.fields.CDN,
				lock:            tt.fields.lock,
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

func TestLocal_Delete(t *testing.T) {
	type fields struct {
		AbstractAdapter gfs.AbstractAdapter
		root            string
		CDN             string
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
			f := &Local{
				AbstractAdapter: tt.fields.AbstractAdapter,
				root:            tt.fields.root,
				CDN:             tt.fields.CDN,
				lock:            tt.fields.lock,
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

func TestLocal_DiskName(t *testing.T) {
	type fields struct {
		AbstractAdapter gfs.AbstractAdapter
		root            string
		CDN             string
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
			f := Local{
				AbstractAdapter: tt.fields.AbstractAdapter,
				root:            tt.fields.root,
				CDN:             tt.fields.CDN,
				lock:            tt.fields.lock,
			}
			if got := f.DiskName(); got != tt.want {
				t.Errorf("DiskName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLocal_Exists(t *testing.T) {
	type fields struct {
		AbstractAdapter gfs.AbstractAdapter
		root            string
		CDN             string
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
			f := &Local{
				AbstractAdapter: tt.fields.AbstractAdapter,
				root:            tt.fields.root,
				CDN:             tt.fields.CDN,
				lock:            tt.fields.lock,
			}
			got, err := f.Exists(tt.args.path)
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

func TestLocal_GFSAdapter(t *testing.T) {
	type fields struct {
		AbstractAdapter gfs.AbstractAdapter
		root            string
		CDN             string
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
			f := &Local{
				AbstractAdapter: tt.fields.AbstractAdapter,
				root:            tt.fields.root,
				CDN:             tt.fields.CDN,
				lock:            tt.fields.lock,
			}
			if got := f.GFSAdapter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GFSAdapter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLocal_MimeType(t *testing.T) {
	type fields struct {
		AbstractAdapter gfs.AbstractAdapter
		root            string
		CDN             string
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
			f := &Local{
				AbstractAdapter: tt.fields.AbstractAdapter,
				root:            tt.fields.root,
				CDN:             tt.fields.CDN,
				lock:            tt.fields.lock,
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

func TestLocal_Move(t *testing.T) {
	type fields struct {
		AbstractAdapter gfs.AbstractAdapter
		root            string
		CDN             string
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
			f := &Local{
				AbstractAdapter: tt.fields.AbstractAdapter,
				root:            tt.fields.root,
				CDN:             tt.fields.CDN,
				lock:            tt.fields.lock,
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

func TestLocal_Read(t *testing.T) {
	type fields struct {
		AbstractAdapter gfs.AbstractAdapter
		root            string
		CDN             string
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
			f := &Local{
				AbstractAdapter: tt.fields.AbstractAdapter,
				root:            tt.fields.root,
				CDN:             tt.fields.CDN,
				lock:            tt.fields.lock,
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

func TestLocal_Size(t *testing.T) {
	type fields struct {
		AbstractAdapter gfs.AbstractAdapter
		root            string
		CDN             string
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
			f := &Local{
				AbstractAdapter: tt.fields.AbstractAdapter,
				root:            tt.fields.root,
				CDN:             tt.fields.CDN,
				lock:            tt.fields.lock,
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

func TestLocal_URL(t *testing.T) {
	type fields struct {
		AbstractAdapter gfs.AbstractAdapter
		root            string
		CDN             string
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
			f := &Local{
				AbstractAdapter: tt.fields.AbstractAdapter,
				root:            tt.fields.root,
				CDN:             tt.fields.CDN,
				lock:            tt.fields.lock,
			}
			got, err := f.URL(tt.args.path)
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

func TestLocal_Update(t *testing.T) {
	type fields struct {
		AbstractAdapter gfs.AbstractAdapter
		root            string
		CDN             string
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
			f := &Local{
				AbstractAdapter: tt.fields.AbstractAdapter,
				root:            tt.fields.root,
				CDN:             tt.fields.CDN,
				lock:            tt.fields.lock,
			}
			got, err := f.Update(tt.args.path, tt.args.contents)
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

func TestLocal_UpdateStream(t *testing.T) {
	type fields struct {
		AbstractAdapter gfs.AbstractAdapter
		root            string
		CDN             string
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
			f := &Local{
				AbstractAdapter: tt.fields.AbstractAdapter,
				root:            tt.fields.root,
				CDN:             tt.fields.CDN,
				lock:            tt.fields.lock,
			}
			got, err := f.UpdateStream(tt.args.path, tt.args.resource)
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

func TestLocal_Write(t *testing.T) {
	type fields struct {
		AbstractAdapter gfs.AbstractAdapter
		root            string
		CDN             string
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
			f := &Local{
				AbstractAdapter: tt.fields.AbstractAdapter,
				root:            tt.fields.root,
				CDN:             tt.fields.CDN,
				lock:            tt.fields.lock,
			}
			got, err := f.Write(tt.args.path, tt.args.contents)
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

func TestLocal_WriteReader(t *testing.T) {
	type fields struct {
		AbstractAdapter gfs.AbstractAdapter
		root            string
		CDN             string
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
			f := &Local{
				AbstractAdapter: tt.fields.AbstractAdapter,
				root:            tt.fields.root,
				CDN:             tt.fields.CDN,
				lock:            tt.fields.lock,
			}
			got, err := f.WriteReader(tt.args.path, tt.args.reader)
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

func TestLocal_WriteStream(t *testing.T) {
	type fields struct {
		AbstractAdapter gfs.AbstractAdapter
		root            string
		CDN             string
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
			f := &Local{
				AbstractAdapter: tt.fields.AbstractAdapter,
				root:            tt.fields.root,
				CDN:             tt.fields.CDN,
				lock:            tt.fields.lock,
			}
			got, err := f.WriteStream(tt.args.path, tt.args.resource)
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

func TestLocal_ensureDirectory(t *testing.T) {
	type fields struct {
		AbstractAdapter gfs.AbstractAdapter
		root            string
		CDN             string
		lock            *sync.Mutex
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
			f := &Local{
				AbstractAdapter: tt.fields.AbstractAdapter,
				root:            tt.fields.root,
				CDN:             tt.fields.CDN,
				lock:            tt.fields.lock,
			}
			if err := f.ensureDirectory(tt.args.root); (err != nil) != tt.wantErr {
				t.Errorf("ensureDirectory() error = %v, wantErr %v", err, tt.wantErr)
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

func TestNewLocal(t *testing.T) {
	type args struct {
		config *Config
	}
	tests := []struct {
		name string
		args args
		want *Local
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
