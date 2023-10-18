package fsoss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/pkg6/go-flysystem/gfs"
	"net/url"
	"reflect"
	"testing"
)

func TestConfig_BucketUrl(t *testing.T) {
	type fields struct {
		CDN             string
		Bucket          string
		Endpoint        string
		AccessKeyID     string
		AccessKeySecret string
		OssConfig       *oss.Config
	}
	tests := []struct {
		name    string
		fields  fields
		want    *url.URL
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				CDN:             tt.fields.CDN,
				Bucket:          tt.fields.Bucket,
				Endpoint:        tt.fields.Endpoint,
				AccessKeyID:     tt.fields.AccessKeyID,
				AccessKeySecret: tt.fields.AccessKeySecret,
				OssConfig:       tt.fields.OssConfig,
			}
			got, err := c.BucketUrl()
			if (err != nil) != tt.wantErr {
				t.Errorf("BucketUrl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BucketUrl() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfig_New(t *testing.T) {
	type fields struct {
		CDN             string
		Bucket          string
		Endpoint        string
		AccessKeyID     string
		AccessKeySecret string
		OssConfig       *oss.Config
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
				CDN:             tt.fields.CDN,
				Bucket:          tt.fields.Bucket,
				Endpoint:        tt.fields.Endpoint,
				AccessKeyID:     tt.fields.AccessKeyID,
				AccessKeySecret: tt.fields.AccessKeySecret,
				OssConfig:       tt.fields.OssConfig,
			}
			if got := c.New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfig_URL(t *testing.T) {
	type fields struct {
		CDN             string
		Bucket          string
		Endpoint        string
		AccessKeyID     string
		AccessKeySecret string
		OssConfig       *oss.Config
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
				CDN:             tt.fields.CDN,
				Bucket:          tt.fields.Bucket,
				Endpoint:        tt.fields.Endpoint,
				AccessKeyID:     tt.fields.AccessKeyID,
				AccessKeySecret: tt.fields.AccessKeySecret,
				OssConfig:       tt.fields.OssConfig,
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
