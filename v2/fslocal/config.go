package fslocal

import (
	"github.com/pkg6/go-flysystem/v2"
)

type Config struct {
}

func (c *Config) New() v2.IAdapter {
	return NewLocal(c)
}
