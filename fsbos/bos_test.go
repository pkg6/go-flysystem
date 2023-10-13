package fsbos

import (
	"github.com/pkg6/go-flysystem"
	"github.com/pkg6/go-flysystem/v2"
	"github.com/pkg6/go-flysystem/v2/fsbos"
	"io"
	"reflect"
	"testing"
)

func TestFSBos_Adapter(t *testing.T) {
	type fields struct {
		AbstractAdapter v2.AbstractAdapter
		Config          *Config
	}
	tests := []struct {
		name   string
		fields fields
		want   *fsbos.Adapter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &FSBos{
				AbstractAdapter: tt.fields.AbstractAdapter,
				Config:          tt.fields.Config,
			}
			if got := f.Adapter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Adapter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFSBos_Copy(t *testing.T) {
	type fields struct {
		AbstractAdapter v2.AbstractAdapter
		Config          *Config
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
			f := &FSBos{
				AbstractAdapter: tt.fields.AbstractAdapter,
				Config:          tt.fields.Config,
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

func TestFSBos_Delete(t *testing.T) {
	type fields struct {
		AbstractAdapter v2.AbstractAdapter
		Config          *Config
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
			f := &FSBos{
				AbstractAdapter: tt.fields.AbstractAdapter,
				Config:          tt.fields.Config,
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

func TestFSBos_DiskName(t *testing.T) {
	type fields struct {
		AbstractAdapter v2.AbstractAdapter
		Config          *Config
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
			f := &FSBos{
				AbstractAdapter: tt.fields.AbstractAdapter,
				Config:          tt.fields.Config,
			}
			if got := f.DiskName(); got != tt.want {
				t.Errorf("DiskName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFSBos_Exists(t *testing.T) {
	type fields struct {
		AbstractAdapter v2.AbstractAdapter
		Config          *Config
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
			f := &FSBos{
				AbstractAdapter: tt.fields.AbstractAdapter,
				Config:          tt.fields.Config,
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

func TestFSBos_MimeType(t *testing.T) {
	type fields struct {
		AbstractAdapter v2.AbstractAdapter
		Config          *Config
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
			f := &FSBos{
				AbstractAdapter: tt.fields.AbstractAdapter,
				Config:          tt.fields.Config,
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

func TestFSBos_Move(t *testing.T) {
	type fields struct {
		AbstractAdapter v2.AbstractAdapter
		Config          *Config
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
			f := &FSBos{
				AbstractAdapter: tt.fields.AbstractAdapter,
				Config:          tt.fields.Config,
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

func TestFSBos_Read(t *testing.T) {
	type fields struct {
		AbstractAdapter v2.AbstractAdapter
		Config          *Config
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
			f := &FSBos{
				AbstractAdapter: tt.fields.AbstractAdapter,
				Config:          tt.fields.Config,
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

func TestFSBos_Size(t *testing.T) {
	type fields struct {
		AbstractAdapter v2.AbstractAdapter
		Config          *Config
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
			f := &FSBos{
				AbstractAdapter: tt.fields.AbstractAdapter,
				Config:          tt.fields.Config,
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

func TestFSBos_Update(t *testing.T) {
	type fields struct {
		AbstractAdapter v2.AbstractAdapter
		Config          *Config
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
			f := &FSBos{
				AbstractAdapter: tt.fields.AbstractAdapter,
				Config:          tt.fields.Config,
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

func TestFSBos_UpdateStream(t *testing.T) {
	type fields struct {
		AbstractAdapter v2.AbstractAdapter
		Config          *Config
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
			f := &FSBos{
				AbstractAdapter: tt.fields.AbstractAdapter,
				Config:          tt.fields.Config,
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

func TestFSBos_Write(t *testing.T) {
	type fields struct {
		AbstractAdapter v2.AbstractAdapter
		Config          *Config
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
			f := &FSBos{
				AbstractAdapter: tt.fields.AbstractAdapter,
				Config:          tt.fields.Config,
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

func TestFSBos_WriteReader(t *testing.T) {
	type fields struct {
		AbstractAdapter v2.AbstractAdapter
		Config          *Config
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
			f := &FSBos{
				AbstractAdapter: tt.fields.AbstractAdapter,
				Config:          tt.fields.Config,
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

func TestFSBos_WriteStream(t *testing.T) {
	type fields struct {
		AbstractAdapter v2.AbstractAdapter
		Config          *Config
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
			f := &FSBos{
				AbstractAdapter: tt.fields.AbstractAdapter,
				Config:          tt.fields.Config,
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
