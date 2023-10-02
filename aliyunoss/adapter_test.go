package aliyunoss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/pkg6/go-flysystem"
	"io"
	"net/http"
	"reflect"
	"sync"
	"testing"
)

func TestAdapter_Clone(t *testing.T) {
	type fields struct {
		AbstractAdapter flysystem.AbstractAdapter
		Config          Config
		Oss             *oss.Client
		lock            *sync.Mutex
	}
	tests := []struct {
		name   string
		fields fields
		want   flysystem.IAdapter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := Adapter{
				AbstractAdapter: tt.fields.AbstractAdapter,
				Config:          tt.fields.Config,
				Oss:             tt.fields.Oss,
				lock:            tt.fields.lock,
			}
			if got := f.Clone(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Clone() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdapter_Copy(t *testing.T) {
	type fields struct {
		AbstractAdapter flysystem.AbstractAdapter
		Config          Config
		Oss             *oss.Client
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
			f := &Adapter{
				AbstractAdapter: tt.fields.AbstractAdapter,
				Config:          tt.fields.Config,
				Oss:             tt.fields.Oss,
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

func TestAdapter_CreateDirectory(t *testing.T) {
	type fields struct {
		AbstractAdapter flysystem.AbstractAdapter
		Config          Config
		Oss             *oss.Client
		lock            *sync.Mutex
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
				AbstractAdapter: tt.fields.AbstractAdapter,
				Config:          tt.fields.Config,
				Oss:             tt.fields.Oss,
				lock:            tt.fields.lock,
			}
			if err := f.CreateDirectory(tt.args.dirname); (err != nil) != tt.wantErr {
				t.Errorf("CreateDirectory() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAdapter_Delete(t *testing.T) {
	type fields struct {
		AbstractAdapter flysystem.AbstractAdapter
		Config          Config
		Oss             *oss.Client
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
			f := &Adapter{
				AbstractAdapter: tt.fields.AbstractAdapter,
				Config:          tt.fields.Config,
				Oss:             tt.fields.Oss,
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

func TestAdapter_DeleteDirectory(t *testing.T) {
	type fields struct {
		AbstractAdapter flysystem.AbstractAdapter
		Config          Config
		Oss             *oss.Client
		lock            *sync.Mutex
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
				AbstractAdapter: tt.fields.AbstractAdapter,
				Config:          tt.fields.Config,
				Oss:             tt.fields.Oss,
				lock:            tt.fields.lock,
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
		AbstractAdapter flysystem.AbstractAdapter
		Config          Config
		Oss             *oss.Client
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
			f := &Adapter{
				AbstractAdapter: tt.fields.AbstractAdapter,
				Config:          tt.fields.Config,
				Oss:             tt.fields.Oss,
				lock:            tt.fields.lock,
			}
			if got := f.DiskName(); got != tt.want {
				t.Errorf("DiskName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdapter_Exists(t *testing.T) {
	type fields struct {
		AbstractAdapter flysystem.AbstractAdapter
		Config          Config
		Oss             *oss.Client
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
			f := &Adapter{
				AbstractAdapter: tt.fields.AbstractAdapter,
				Config:          tt.fields.Config,
				Oss:             tt.fields.Oss,
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

func TestAdapter_MimeType(t *testing.T) {
	type fields struct {
		AbstractAdapter flysystem.AbstractAdapter
		Config          Config
		Oss             *oss.Client
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
			f := &Adapter{
				AbstractAdapter: tt.fields.AbstractAdapter,
				Config:          tt.fields.Config,
				Oss:             tt.fields.Oss,
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

func TestAdapter_Move(t *testing.T) {
	type fields struct {
		AbstractAdapter flysystem.AbstractAdapter
		Config          Config
		Oss             *oss.Client
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
			f := &Adapter{
				AbstractAdapter: tt.fields.AbstractAdapter,
				Config:          tt.fields.Config,
				Oss:             tt.fields.Oss,
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

func TestAdapter_Read(t *testing.T) {
	type fields struct {
		AbstractAdapter flysystem.AbstractAdapter
		Config          Config
		Oss             *oss.Client
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
			f := &Adapter{
				AbstractAdapter: tt.fields.AbstractAdapter,
				Config:          tt.fields.Config,
				Oss:             tt.fields.Oss,
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

func TestAdapter_Size(t *testing.T) {
	type fields struct {
		AbstractAdapter flysystem.AbstractAdapter
		Config          Config
		Oss             *oss.Client
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
			f := &Adapter{
				AbstractAdapter: tt.fields.AbstractAdapter,
				Config:          tt.fields.Config,
				Oss:             tt.fields.Oss,
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

func TestAdapter_Update(t *testing.T) {
	type fields struct {
		AbstractAdapter flysystem.AbstractAdapter
		Config          Config
		Oss             *oss.Client
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
			f := &Adapter{
				AbstractAdapter: tt.fields.AbstractAdapter,
				Config:          tt.fields.Config,
				Oss:             tt.fields.Oss,
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

func TestAdapter_UpdateStream(t *testing.T) {
	type fields struct {
		AbstractAdapter flysystem.AbstractAdapter
		Config          Config
		Oss             *oss.Client
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
			f := &Adapter{
				AbstractAdapter: tt.fields.AbstractAdapter,
				Config:          tt.fields.Config,
				Oss:             tt.fields.Oss,
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

func TestAdapter_Write(t *testing.T) {
	type fields struct {
		AbstractAdapter flysystem.AbstractAdapter
		Config          Config
		Oss             *oss.Client
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
			f := &Adapter{
				AbstractAdapter: tt.fields.AbstractAdapter,
				Config:          tt.fields.Config,
				Oss:             tt.fields.Oss,
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

func TestAdapter_WriteReader(t *testing.T) {
	type fields struct {
		AbstractAdapter flysystem.AbstractAdapter
		Config          Config
		Oss             *oss.Client
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
			f := &Adapter{
				AbstractAdapter: tt.fields.AbstractAdapter,
				Config:          tt.fields.Config,
				Oss:             tt.fields.Oss,
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

func TestAdapter_WriteStream(t *testing.T) {
	type fields struct {
		AbstractAdapter flysystem.AbstractAdapter
		Config          Config
		Oss             *oss.Client
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
			f := &Adapter{
				AbstractAdapter: tt.fields.AbstractAdapter,
				Config:          tt.fields.Config,
				Oss:             tt.fields.Oss,
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

func TestAdapter_copyObject(t *testing.T) {
	type fields struct {
		AbstractAdapter flysystem.AbstractAdapter
		Config          Config
		Oss             *oss.Client
		lock            *sync.Mutex
	}
	type args struct {
		srcObjectKey  string
		destObjectKey string
		isDelete      bool
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
				AbstractAdapter: tt.fields.AbstractAdapter,
				Config:          tt.fields.Config,
				Oss:             tt.fields.Oss,
				lock:            tt.fields.lock,
			}
			got, err := f.copyObject(tt.args.srcObjectKey, tt.args.destObjectKey, tt.args.isDelete)
			if (err != nil) != tt.wantErr {
				t.Errorf("copyObject() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("copyObject() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdapter_getObjectMeta(t *testing.T) {
	type fields struct {
		AbstractAdapter flysystem.AbstractAdapter
		Config          Config
		Oss             *oss.Client
		lock            *sync.Mutex
	}
	type args struct {
		path string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantHeader http.Header
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Adapter{
				AbstractAdapter: tt.fields.AbstractAdapter,
				Config:          tt.fields.Config,
				Oss:             tt.fields.Oss,
				lock:            tt.fields.lock,
			}
			gotHeader, err := f.getObjectMeta(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("getObjectMeta() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotHeader, tt.wantHeader) {
				t.Errorf("getObjectMeta() gotHeader = %v, want %v", gotHeader, tt.wantHeader)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		config Config
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
