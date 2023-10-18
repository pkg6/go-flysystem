package fsbos

import (
	"github.com/pkg6/go-flysystem/gfs"
	"net/url"
	"reflect"
	"testing"
)

func TestConfig_BucketUrl(t *testing.T) {
	type fields struct {
		CDN              string
		Ak               string
		Sk               string
		Endpoint         string
		RedirectDisabled bool
		Bucket           string
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
				CDN:              tt.fields.CDN,
				Ak:               tt.fields.Ak,
				Sk:               tt.fields.Sk,
				Endpoint:         tt.fields.Endpoint,
				RedirectDisabled: tt.fields.RedirectDisabled,
				Bucket:           tt.fields.Bucket,
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
		CDN              string
		Ak               string
		Sk               string
		Endpoint         string
		RedirectDisabled bool
		Bucket           string
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
				CDN:              tt.fields.CDN,
				Ak:               tt.fields.Ak,
				Sk:               tt.fields.Sk,
				Endpoint:         tt.fields.Endpoint,
				RedirectDisabled: tt.fields.RedirectDisabled,
				Bucket:           tt.fields.Bucket,
			}
			if got := c.New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfig_URL(t *testing.T) {
	type fields struct {
		CDN              string
		Ak               string
		Sk               string
		Endpoint         string
		RedirectDisabled bool
		Bucket           string
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
				CDN:              tt.fields.CDN,
				Ak:               tt.fields.Ak,
				Sk:               tt.fields.Sk,
				Endpoint:         tt.fields.Endpoint,
				RedirectDisabled: tt.fields.RedirectDisabled,
				Bucket:           tt.fields.Bucket,
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
