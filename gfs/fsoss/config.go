package fsoss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/pkg6/go-flysystem/gfs"
	"net/url"
	"strings"
)

type Config struct {
	CDN             string
	Bucket          string
	Endpoint        string
	AccessKeyID     string
	AccessKeySecret string
	OssConfig       *oss.Config
}

func (c *Config) New() gfs.IAdapter {
	return NewOSS(c)
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
