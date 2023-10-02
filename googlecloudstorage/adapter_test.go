package googlecloudstorage

import (
	"cloud.google.com/go/storage"
	"context"
	"github.com/pkg6/go-flysystem"
	"io"
	"reflect"
	"sync"
	"testing"
)

func TestAdapter_Clone(t *testing.T) {
	type fields struct {
		Config *Config
		ctx    context.Context
		lock   *sync.Mutex
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
			a := &Adapter{
				Config: tt.fields.Config,
				ctx:    tt.fields.ctx,
				lock:   tt.fields.lock,
			}
			if got := a.Clone(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Clone() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdapter_Copy(t *testing.T) {
	type fields struct {
		Config *Config
		ctx    context.Context
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
				ctx:    tt.fields.ctx,
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

func TestAdapter_CopyObject(t *testing.T) {
	type fields struct {
		Config *Config
		ctx    context.Context
		lock   *sync.Mutex
	}
	type args struct {
		source       string
		destination  string
		deleteSource bool
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
				ctx:    tt.fields.ctx,
				lock:   tt.fields.lock,
			}
			got, err := a.CopyObject(tt.args.source, tt.args.destination, tt.args.deleteSource)
			if (err != nil) != tt.wantErr {
				t.Errorf("CopyObject() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CopyObject() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdapter_Delete(t *testing.T) {
	type fields struct {
		Config *Config
		ctx    context.Context
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
				ctx:    tt.fields.ctx,
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
		ctx    context.Context
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
				ctx:    tt.fields.ctx,
				lock:   tt.fields.lock,
			}
			if got := a.DiskName(); got != tt.want {
				t.Errorf("DiskName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdapter_Exists(t *testing.T) {
	type fields struct {
		Config *Config
		ctx    context.Context
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
				ctx:    tt.fields.ctx,
				lock:   tt.fields.lock,
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

func TestAdapter_GetMetadata(t *testing.T) {
	type fields struct {
		Config *Config
		ctx    context.Context
		lock   *sync.Mutex
	}
	type args struct {
		object string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *storage.ObjectAttrs
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Adapter{
				Config: tt.fields.Config,
				ctx:    tt.fields.ctx,
				lock:   tt.fields.lock,
			}
			got, err := a.GetMetadata(tt.args.object)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMetadata() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMetadata() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdapter_MimeType(t *testing.T) {
	type fields struct {
		Config *Config
		ctx    context.Context
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
				ctx:    tt.fields.ctx,
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
		ctx    context.Context
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
				ctx:    tt.fields.ctx,
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
		ctx    context.Context
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
				ctx:    tt.fields.ctx,
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
		ctx    context.Context
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
				ctx:    tt.fields.ctx,
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

func TestAdapter_StorageClient(t *testing.T) {
	type fields struct {
		Config *Config
		ctx    context.Context
		lock   *sync.Mutex
	}
	tests := []struct {
		name    string
		fields  fields
		want    *storage.Client
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Adapter{
				Config: tt.fields.Config,
				ctx:    tt.fields.ctx,
				lock:   tt.fields.lock,
			}
			got, err := a.StorageClient()
			if (err != nil) != tt.wantErr {
				t.Errorf("StorageClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StorageClient() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdapter_StorageObject(t *testing.T) {
	type fields struct {
		Config *Config
		ctx    context.Context
		lock   *sync.Mutex
	}
	type args struct {
		object string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *storage.ObjectHandle
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Adapter{
				Config: tt.fields.Config,
				ctx:    tt.fields.ctx,
				lock:   tt.fields.lock,
			}
			got, err := a.StorageObject(tt.args.object)
			if (err != nil) != tt.wantErr {
				t.Errorf("StorageObject() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StorageObject() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdapter_StorageObjectTimeout(t *testing.T) {
	type fields struct {
		Config *Config
		ctx    context.Context
		lock   *sync.Mutex
	}
	type args struct {
		object string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *storage.ObjectHandle
		want1   context.Context
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Adapter{
				Config: tt.fields.Config,
				ctx:    tt.fields.ctx,
				lock:   tt.fields.lock,
			}
			got, got1, err := a.StorageObjectTimeout(tt.args.object)
			if (err != nil) != tt.wantErr {
				t.Errorf("StorageObjectTimeout() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StorageObjectTimeout() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("StorageObjectTimeout() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestAdapter_Update(t *testing.T) {
	type fields struct {
		Config *Config
		ctx    context.Context
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
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Adapter{
				Config: tt.fields.Config,
				ctx:    tt.fields.ctx,
				lock:   tt.fields.lock,
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

func TestAdapter_UpdateStream(t *testing.T) {
	type fields struct {
		Config *Config
		ctx    context.Context
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
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Adapter{
				Config: tt.fields.Config,
				ctx:    tt.fields.ctx,
				lock:   tt.fields.lock,
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

func TestAdapter_Write(t *testing.T) {
	type fields struct {
		Config *Config
		ctx    context.Context
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
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Adapter{
				Config: tt.fields.Config,
				ctx:    tt.fields.ctx,
				lock:   tt.fields.lock,
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

func TestAdapter_WriteReader(t *testing.T) {
	type fields struct {
		Config *Config
		ctx    context.Context
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
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Adapter{
				Config: tt.fields.Config,
				ctx:    tt.fields.ctx,
				lock:   tt.fields.lock,
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

func TestAdapter_WriteStream(t *testing.T) {
	type fields struct {
		Config *Config
		ctx    context.Context
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
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Adapter{
				Config: tt.fields.Config,
				ctx:    tt.fields.ctx,
				lock:   tt.fields.lock,
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
