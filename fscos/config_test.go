package fscos

import (
	"github.com/pkg6/go-flysystem"
	"reflect"
	"testing"
)

func TestConfig_New(t *testing.T) {
	type fields struct {
		CDN        string
		BucketURL  string
		SecretID   string
		SecretKey  string
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
				BucketURL:  tt.fields.BucketURL,
				SecretID:   tt.fields.SecretID,
				SecretKey:  tt.fields.SecretKey,
				PathPrefix: tt.fields.PathPrefix,
			}
			if got := c.New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
