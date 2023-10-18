package gfs

import (
	"io"
	"net/url"
	"reflect"
	"sync"
	"testing"
)

func TestFsManage_Copy(t *testing.T) {
	type fields struct {
		disk         string
		disks        []string
		diskAdapters map[string]IAdapter
		l            *sync.Mutex
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
			f := &FsManage{
				disk:         tt.fields.disk,
				disks:        tt.fields.disks,
				diskAdapters: tt.fields.diskAdapters,
				l:            tt.fields.l,
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

func TestFsManage_Delete(t *testing.T) {
	type fields struct {
		disk         string
		disks        []string
		diskAdapters map[string]IAdapter
		l            *sync.Mutex
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
			f := &FsManage{
				disk:         tt.fields.disk,
				disks:        tt.fields.disks,
				diskAdapters: tt.fields.diskAdapters,
				l:            tt.fields.l,
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

func TestFsManage_Disk(t *testing.T) {
	type fields struct {
		disk         string
		disks        []string
		diskAdapters map[string]IAdapter
		l            *sync.Mutex
	}
	type args struct {
		disk string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   IFSManage
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &FsManage{
				disk:         tt.fields.disk,
				disks:        tt.fields.disks,
				diskAdapters: tt.fields.diskAdapters,
				l:            tt.fields.l,
			}
			if got := f.Disk(tt.args.disk); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Disk() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFsManage_Exist(t *testing.T) {
	type fields struct {
		disk         string
		disks        []string
		diskAdapters map[string]IAdapter
		l            *sync.Mutex
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
			f := &FsManage{
				disk:         tt.fields.disk,
				disks:        tt.fields.disks,
				diskAdapters: tt.fields.diskAdapters,
				l:            tt.fields.l,
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

func TestFsManage_Extend(t *testing.T) {
	type fields struct {
		disk         string
		disks        []string
		diskAdapters map[string]IAdapter
		l            *sync.Mutex
	}
	type args struct {
		adapter IAdapter
		names   []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   IFSManage
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &FsManage{
				disk:         tt.fields.disk,
				disks:        tt.fields.disks,
				diskAdapters: tt.fields.diskAdapters,
				l:            tt.fields.l,
			}
			if got := f.Extend(tt.args.adapter, tt.args.names...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Extend() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFsManage_FindAdapter(t *testing.T) {
	type fields struct {
		disk         string
		disks        []string
		diskAdapters map[string]IAdapter
		l            *sync.Mutex
	}
	tests := []struct {
		name   string
		fields fields
		want   IAdapter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &FsManage{
				disk:         tt.fields.disk,
				disks:        tt.fields.disks,
				diskAdapters: tt.fields.diskAdapters,
				l:            tt.fields.l,
			}
			if got := f.FindAdapter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindAdapter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFsManage_MimeType(t *testing.T) {
	type fields struct {
		disk         string
		disks        []string
		diskAdapters map[string]IAdapter
		l            *sync.Mutex
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
			f := &FsManage{
				disk:         tt.fields.disk,
				disks:        tt.fields.disks,
				diskAdapters: tt.fields.diskAdapters,
				l:            tt.fields.l,
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

func TestFsManage_Move(t *testing.T) {
	type fields struct {
		disk         string
		disks        []string
		diskAdapters map[string]IAdapter
		l            *sync.Mutex
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
			f := &FsManage{
				disk:         tt.fields.disk,
				disks:        tt.fields.disks,
				diskAdapters: tt.fields.diskAdapters,
				l:            tt.fields.l,
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

func TestFsManage_Read(t *testing.T) {
	type fields struct {
		disk         string
		disks        []string
		diskAdapters map[string]IAdapter
		l            *sync.Mutex
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
			f := &FsManage{
				disk:         tt.fields.disk,
				disks:        tt.fields.disks,
				diskAdapters: tt.fields.diskAdapters,
				l:            tt.fields.l,
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

func TestFsManage_Size(t *testing.T) {
	type fields struct {
		disk         string
		disks        []string
		diskAdapters map[string]IAdapter
		l            *sync.Mutex
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
			f := &FsManage{
				disk:         tt.fields.disk,
				disks:        tt.fields.disks,
				diskAdapters: tt.fields.diskAdapters,
				l:            tt.fields.l,
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

func TestFsManage_URL(t *testing.T) {
	type fields struct {
		disk         string
		disks        []string
		diskAdapters map[string]IAdapter
		l            *sync.Mutex
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
			f := &FsManage{
				disk:         tt.fields.disk,
				disks:        tt.fields.disks,
				diskAdapters: tt.fields.diskAdapters,
				l:            tt.fields.l,
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

func TestFsManage_Update(t *testing.T) {
	type fields struct {
		disk         string
		disks        []string
		diskAdapters map[string]IAdapter
		l            *sync.Mutex
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
			f := &FsManage{
				disk:         tt.fields.disk,
				disks:        tt.fields.disks,
				diskAdapters: tt.fields.diskAdapters,
				l:            tt.fields.l,
			}
			if err := f.Update(tt.args.path, tt.args.contents); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFsManage_UpdateStream(t *testing.T) {
	type fields struct {
		disk         string
		disks        []string
		diskAdapters map[string]IAdapter
		l            *sync.Mutex
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
			f := &FsManage{
				disk:         tt.fields.disk,
				disks:        tt.fields.disks,
				diskAdapters: tt.fields.diskAdapters,
				l:            tt.fields.l,
			}
			if err := f.UpdateStream(tt.args.path, tt.args.resource); (err != nil) != tt.wantErr {
				t.Errorf("UpdateStream() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFsManage_Write(t *testing.T) {
	type fields struct {
		disk         string
		disks        []string
		diskAdapters map[string]IAdapter
		l            *sync.Mutex
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
			f := &FsManage{
				disk:         tt.fields.disk,
				disks:        tt.fields.disks,
				diskAdapters: tt.fields.diskAdapters,
				l:            tt.fields.l,
			}
			if err := f.Write(tt.args.path, tt.args.contents); (err != nil) != tt.wantErr {
				t.Errorf("Write() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFsManage_WriteReader(t *testing.T) {
	type fields struct {
		disk         string
		disks        []string
		diskAdapters map[string]IAdapter
		l            *sync.Mutex
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
			f := &FsManage{
				disk:         tt.fields.disk,
				disks:        tt.fields.disks,
				diskAdapters: tt.fields.diskAdapters,
				l:            tt.fields.l,
			}
			if err := f.WriteReader(tt.args.path, tt.args.reader); (err != nil) != tt.wantErr {
				t.Errorf("WriteReader() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFsManage_WriteStream(t *testing.T) {
	type fields struct {
		disk         string
		disks        []string
		diskAdapters map[string]IAdapter
		l            *sync.Mutex
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
			f := &FsManage{
				disk:         tt.fields.disk,
				disks:        tt.fields.disks,
				diskAdapters: tt.fields.diskAdapters,
				l:            tt.fields.l,
			}
			if err := f.WriteStream(tt.args.path, tt.args.resource); (err != nil) != tt.wantErr {
				t.Errorf("WriteStream() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want IFSManage
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewConfig(t *testing.T) {
	type args struct {
		config IConfig
	}
	tests := []struct {
		name string
		args args
		want IFSManage
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewConfig(tt.args.config); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
