package local

import "github.com/pkg6/go-flysystem"

type Config struct {
	Root string
	CDN  string
}

func (c *Config) New() flysystem.IAdapter {
	return NewLocal(c)
}
