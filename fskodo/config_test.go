package fskodo

import (
	"github.com/pkg6/go-flysystem"
	"github.com/qiniu/go-sdk/v7/storage"
	"reflect"
	"testing"
)

func TestConfig_New(t *testing.T) {
	type fields struct {
		CDN        string
		AccessKey  string
		SecretKey  string
		Bucket     string
		Policy     *storage.PutPolicy
		Config     *storage.Config
		PathPrefix string
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
			c := &Config{
				CDN:        tt.fields.CDN,
				AccessKey:  tt.fields.AccessKey,
				SecretKey:  tt.fields.SecretKey,
				Bucket:     tt.fields.Bucket,
				Policy:     tt.fields.Policy,
				Config:     tt.fields.Config,
				PathPrefix: tt.fields.PathPrefix,
			}
			if got := c.New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
