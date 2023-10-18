package fskodo

import (
	v2 "github.com/pkg6/go-flysystem/v2"
	"github.com/qiniu/go-sdk/v7/storage"
	"net/url"
)

type Config struct {
	CDN                  string
	AccessKey, SecretKey string
	Bucket               string

	Policy *storage.PutPolicy
	Config *storage.Config
}

func (c *Config) New() v2.IAdapter {
	return NewKoDo(c)
}

func (c *Config) URL(path string) (*url.URL, error) {
	if len(path) > 0 && path[0:1] != "/" {
		path = "/" + path
	}
	return url.Parse(c.CDN + path)
}
