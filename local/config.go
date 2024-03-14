package local

import "github.com/pkg6/go-flysystem"

type Config struct {
	Root string `json:"root" xml:"Root" yaml:"Root"`
	CDN  string `json:"cdn" xml:"CDN" yaml:"CDN"`
}

func (c *Config) New() flysystem.IAdapter {
	return NewLocal(c)
}
