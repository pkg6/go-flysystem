package fsbos

import (
	"github.com/pkg6/go-flysystem/v2"
	"net/url"
	"strings"
)

type Config struct {
	Ak               string
	Sk               string
	Endpoint         string
	RedirectDisabled bool
	Bucket           string
}

func (c *Config) New() v2.IAdapter {
	return NewBOS(c)
}

func (c *Config) URI(path string) (*url.URL, error) {
	endpoint := c.Endpoint
	if !strings.HasPrefix(endpoint, "http") {
		endpoint = "https://" + endpoint
	}
	uri, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}
	return uri.Parse(uri.Scheme + "://" + c.Bucket + "." + uri.Host + "/" + path)
}
