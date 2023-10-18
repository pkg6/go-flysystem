package fslocal

import (
	v2 "github.com/pkg6/go-flysystem/v2"
	"net/url"
)

type Config struct {
	CDN string
}

func (c *Config) New() v2.IAdapter {
	return NewLocal(c)
}
func (c *Config) URL(path string) (*url.URL, error) {
	if len(path) > 0 && path[0:1] != "/" {
		path = "/" + path
	}
	return url.Parse(c.CDN + path)
}
