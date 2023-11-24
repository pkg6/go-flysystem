package fsoss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/pkg6/go-flysystem"
)

type Config struct {
	CDN             string
	Bucket          string
	Endpoint        string
	AccessKeyID     string
	AccessKeySecret string
	OssConfig       *oss.Config
	PathPrefix      string
}

func (c *Config) New() flysystem.IAdapter {
	return NewOSS(c)
}
