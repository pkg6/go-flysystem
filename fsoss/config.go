package fsoss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/pkg6/go-flysystem"
)

type Config struct {
	CDN    string `json:"cdn" xml:"CDN" yaml:"CDN"`
	Bucket string `json:"bucket" xml:"Bucket" yaml:"Bucket"`
	//https://help.aliyun.com/zh/oss/user-guide/regions-and-endpoints
	Endpoint        string      `json:"endpoint" xml:"Endpoint" yaml:"Endpoint"`
	AccessKeyID     string      `json:"access_key_id" xml:"AccessKeyID" yaml:"AccessKeyID"`
	AccessKeySecret string      `json:"access_key_secret" xml:"AccessKeySecret" yaml:"AccessKeySecret"`
	Config          *oss.Config `json:"config" xml:"Config" yaml:"Config"`
	PathPrefix      string      `json:"path_prefix" xml:"PathPrefix" yaml:"PathPrefix"`
}

func (c *Config) New() flysystem.IAdapter {
	return NewOSS(c)
}
