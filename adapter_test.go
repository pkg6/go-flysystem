package flysystem

import "testing"

func TestAbstractAdapter_ApplyPathPrefix(t *testing.T) {
	type fields struct {
		prefix string
	}
	type args struct {
		path string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AbstractAdapter{
				prefix: tt.fields.prefix,
			}
			if got := a.ApplyPathPrefix(tt.args.path); got != tt.want {
				t.Errorf("ApplyPathPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAbstractAdapter_SetPathPrefix(t *testing.T) {
	type fields struct {
		prefix string
	}
	type args struct {
		prefix string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AbstractAdapter{
				prefix: tt.fields.prefix,
			}
			a.SetPathPrefix(tt.args.prefix)
		})
	}
}
