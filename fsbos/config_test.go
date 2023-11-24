package fsbos

import (
	"github.com/pkg6/go-flysystem"
	"reflect"
	"testing"
)

func TestConfig_New(t *testing.T) {
	type fields struct {
		CDN              string
		Ak               string
		Sk               string
		Endpoint         string
		RedirectDisabled bool
		Bucket           string
		PathPrefix       string
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
				CDN:              tt.fields.CDN,
				Ak:               tt.fields.Ak,
				Sk:               tt.fields.Sk,
				Endpoint:         tt.fields.Endpoint,
				RedirectDisabled: tt.fields.RedirectDisabled,
				Bucket:           tt.fields.Bucket,
				PathPrefix:       tt.fields.PathPrefix,
			}
			if got := c.New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
