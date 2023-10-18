package fsbos

import (
	v2 "github.com/pkg6/go-flysystem/v2"
	"net/url"
	"strings"
)

type Config struct {
	CDN              string
	Ak               string
	Sk               string
	Endpoint         string
	RedirectDisabled bool
	Bucket           string
}

func (c *Config) New() v2.IAdapter {
	return NewBOS(c)
}

func (c *Config) URL(path string) (*url.URL, error) {
	bucketUrl, err := c.BucketUrl()
	if err != nil {
		return nil, err
	}
	if len(path) > 0 && path[0:1] != "/" {
		path = "/" + path
	}
	return url.Parse(bucketUrl.String() + path)
}

func (c *Config) BucketUrl() (*url.URL, error) {
	if c.CDN == "" {
		endpoint := c.Endpoint
		if !strings.HasPrefix(endpoint, "http") {
			endpoint = "https://" + endpoint
		}
		uri, _ := url.Parse(endpoint)
		endpointURL, err := uri.Parse(uri.Scheme + "://" + c.Bucket + "." + uri.Host)
		if err != nil {
			return nil, err
		}
		c.CDN = endpointURL.String()
	}
	return url.Parse(c.CDN)
}
