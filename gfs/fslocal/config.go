package fslocal

import (
	"github.com/pkg6/go-flysystem/gfs"
	"net/url"
)

type Config struct {
	CDN string
}

func (c *Config) New() gfs.IAdapter {
	return NewLocal(c)
}
func (c *Config) URL(path string) (*url.URL, error) {
	if len(path) > 0 && path[0:1] != "/" {
		path = "/" + path
	}
	return url.Parse(c.CDN + path)
}
