package fsoss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/pkg6/go-flysystem/v2"
)

type Config struct {
	Bucket          string
	Endpoint        string
	AccessKeyID     string
	AccessKeySecret string
	OssConfig       *oss.Config
}

func (c *Config) New() v2.IAdapter {
	return NewOSS(c)
}
