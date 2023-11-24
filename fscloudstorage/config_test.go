package fscloudstorage

import (
	"github.com/pkg6/go-flysystem"
	"google.golang.org/api/option"
	"reflect"
	"testing"
	"time"
)

func TestConfig_New(t *testing.T) {
	type fields struct {
		CDN             string
		Bucket          string
		WithTimeout     time.Duration
		CredentialsFile string
		Option          []option.ClientOption
		PathPrefix      string
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
				CDN:             tt.fields.CDN,
				Bucket:          tt.fields.Bucket,
				WithTimeout:     tt.fields.WithTimeout,
				CredentialsFile: tt.fields.CredentialsFile,
				Option:          tt.fields.Option,
				PathPrefix:      tt.fields.PathPrefix,
			}
			if got := c.New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
