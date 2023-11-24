package local

import (
	"github.com/pkg6/go-flysystem"
	"reflect"
	"testing"
)

func TestConfig_New(t *testing.T) {
	type fields struct {
		root string
		CDN  string
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
				root: tt.fields.root,
				CDN:  tt.fields.CDN,
			}
			if got := c.New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
