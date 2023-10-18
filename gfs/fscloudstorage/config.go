package fscloudstorage

import (
	"github.com/pkg6/go-flysystem/gfs"
	"google.golang.org/api/option"
	"net/url"
	"time"
)

var (
	DefaultWithTimeout = time.Second * 50
)

type Config struct {
	CDN             string
	Bucket          string
	WithTimeout     time.Duration
	CredentialsFile string
	Option          []option.ClientOption
}

func (c *Config) New() gfs.IAdapter {
	return NewGCS(c)
}

func (c *Config) URL(path string) (*url.URL, error) {
	if len(path) > 0 && path[0:1] != "/" {
		path = "/" + path
	}
	return url.Parse(c.CDN + path)
}
