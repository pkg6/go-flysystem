package fskodo

import (
	"github.com/qiniu/go-sdk/v7/storage"
	"github.com/pkg6/go-flysystem/gfs"
	"net/url"
	"reflect"
	"testing"
)

func TestConfig_New(t *testing.T) {
	type fields struct {
		CDN       string
		AccessKey string
		SecretKey string
		Bucket    string
		Policy    *storage.PutPolicy
		Config    *storage.Config
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
			c := &Config{
				CDN:       tt.fields.CDN,
				AccessKey: tt.fields.AccessKey,
				SecretKey: tt.fields.SecretKey,
				Bucket:    tt.fields.Bucket,
				Policy:    tt.fields.Policy,
				Config:    tt.fields.Config,
			}
			if got := c.New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfig_URL(t *testing.T) {
	type fields struct {
		CDN       string
		AccessKey string
		SecretKey string
		Bucket    string
		Policy    *storage.PutPolicy
		Config    *storage.Config
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
			c := &Config{
				CDN:       tt.fields.CDN,
				AccessKey: tt.fields.AccessKey,
				SecretKey: tt.fields.SecretKey,
				Bucket:    tt.fields.Bucket,
				Policy:    tt.fields.Policy,
				Config:    tt.fields.Config,
			}
			got, err := c.URL(tt.args.path)
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
