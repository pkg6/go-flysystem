package flysystem

import (
	"reflect"
	"testing"
)

func TestFlysystem_Copy(t *testing.T) {
	type fields struct {
		disk         string
		diskAdapters map[string]IAdapter
		diskNames    []string
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
			f := &Flysystem{
				disk:         tt.fields.disk,
				diskAdapters: tt.fields.diskAdapters,
				diskNames:    tt.fields.diskNames,
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

func TestFlysystem_CreateDirectory(t *testing.T) {
	type fields struct {
		disk         string
		diskAdapters map[string]IAdapter
		diskNames    []string
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
			f := &Flysystem{
				disk:         tt.fields.disk,
				diskAdapters: tt.fields.diskAdapters,
				diskNames:    tt.fields.diskNames,
			}
			if err := f.CreateDirectory(tt.args.dirname); (err != nil) != tt.wantErr {
				t.Errorf("CreateDirectory() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFlysystem_Delete(t *testing.T) {
	type fields struct {
		disk         string
		diskAdapters map[string]IAdapter
		diskNames    []string
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
			f := &Flysystem{
				disk:         tt.fields.disk,
				diskAdapters: tt.fields.diskAdapters,
				diskNames:    tt.fields.diskNames,
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

func TestFlysystem_DeleteDirectory(t *testing.T) {
	type fields struct {
		disk         string
		diskAdapters map[string]IAdapter
		diskNames    []string
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
			f := &Flysystem{
				disk:         tt.fields.disk,
				diskAdapters: tt.fields.diskAdapters,
				diskNames:    tt.fields.diskNames,
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

func TestFlysystem_Disk(t *testing.T) {
	type fields struct {
		disk         string
		diskAdapters map[string]IAdapter
		diskNames    []string
	}
	type args struct {
		disk string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   IFlysystem
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Flysystem{
				disk:         tt.fields.disk,
				diskAdapters: tt.fields.diskAdapters,
				diskNames:    tt.fields.diskNames,
			}
			if got := f.Disk(tt.args.disk); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Disk() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlysystem_Exists(t *testing.T) {
	type fields struct {
		disk         string
		diskAdapters map[string]IAdapter
		diskNames    []string
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
			f := &Flysystem{
				disk:         tt.fields.disk,
				diskAdapters: tt.fields.diskAdapters,
				diskNames:    tt.fields.diskNames,
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

func TestFlysystem_Extend(t *testing.T) {
	type fields struct {
		disk         string
		diskAdapters map[string]IAdapter
		diskNames    []string
	}
	type args struct {
		adapter IAdapter
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   IFlysystem
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Flysystem{
				disk:         tt.fields.disk,
				diskAdapters: tt.fields.diskAdapters,
				diskNames:    tt.fields.diskNames,
			}
			if got := f.Extend(tt.args.adapter); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Extend() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlysystem_FindAdapter(t *testing.T) {
	type fields struct {
		disk         string
		diskAdapters map[string]IAdapter
		diskNames    []string
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
			f := &Flysystem{
				disk:         tt.fields.disk,
				diskAdapters: tt.fields.diskAdapters,
				diskNames:    tt.fields.diskNames,
			}
			if got := f.FindAdapter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindAdapter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlysystem_MimeType(t *testing.T) {
	type fields struct {
		disk         string
		diskAdapters map[string]IAdapter
		diskNames    []string
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
			f := &Flysystem{
				disk:         tt.fields.disk,
				diskAdapters: tt.fields.diskAdapters,
				diskNames:    tt.fields.diskNames,
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

func TestFlysystem_Move(t *testing.T) {
	type fields struct {
		disk         string
		diskAdapters map[string]IAdapter
		diskNames    []string
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
			f := &Flysystem{
				disk:         tt.fields.disk,
				diskAdapters: tt.fields.diskAdapters,
				diskNames:    tt.fields.diskNames,
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

func TestFlysystem_Read(t *testing.T) {
	type fields struct {
		disk         string
		diskAdapters map[string]IAdapter
		diskNames    []string
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
			f := &Flysystem{
				disk:         tt.fields.disk,
				diskAdapters: tt.fields.diskAdapters,
				diskNames:    tt.fields.diskNames,
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

func TestFlysystem_Size(t *testing.T) {
	type fields struct {
		disk         string
		diskAdapters map[string]IAdapter
		diskNames    []string
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
			f := &Flysystem{
				disk:         tt.fields.disk,
				diskAdapters: tt.fields.diskAdapters,
				diskNames:    tt.fields.diskNames,
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

func TestFlysystem_Update(t *testing.T) {
	type fields struct {
		disk         string
		diskAdapters map[string]IAdapter
		diskNames    []string
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
			f := &Flysystem{
				disk:         tt.fields.disk,
				diskAdapters: tt.fields.diskAdapters,
				diskNames:    tt.fields.diskNames,
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

func TestFlysystem_UpdateStream(t *testing.T) {
	type fields struct {
		disk         string
		diskAdapters map[string]IAdapter
		diskNames    []string
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
			f := &Flysystem{
				disk:         tt.fields.disk,
				diskAdapters: tt.fields.diskAdapters,
				diskNames:    tt.fields.diskNames,
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

func TestFlysystem_Write(t *testing.T) {
	type fields struct {
		disk         string
		diskAdapters map[string]IAdapter
		diskNames    []string
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
			f := &Flysystem{
				disk:         tt.fields.disk,
				diskAdapters: tt.fields.diskAdapters,
				diskNames:    tt.fields.diskNames,
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

func TestFlysystem_WriteStream(t *testing.T) {
	type fields struct {
		disk         string
		diskAdapters map[string]IAdapter
		diskNames    []string
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
			f := &Flysystem{
				disk:         tt.fields.disk,
				diskAdapters: tt.fields.diskAdapters,
				diskNames:    tt.fields.diskNames,
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
	tests := []struct {
		name string
		want IFlysystem
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

func TestNewAdapters(t *testing.T) {
	type args struct {
		adapters []IAdapter
	}
	tests := []struct {
		name string
		args args
		want IFlysystem
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAdapters(tt.args.adapters...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAdapters() = %v, want %v", got, tt.want)
			}
		})
	}
}
