package fscos

import (
	v2 "github.com/pkg6/go-flysystem/v2"
	"net/url"
)

type Config struct {
	CDN string
	// 存储桶名称，由 bucketname-appid 组成，appid 必须填入，可以在 COS 控制台查看存储桶名称。 https://console.cloud.tencent.com/cos5/bucket
	// 替换为用户的 region，存储桶 region 可以在 COS 控制台“存储桶概览”查看 https://console.cloud.tencent.com/
	BucketURL string
	// 用户的 SecretId，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参见 https://cloud.tencent.com/document/product/598/37140
	// 用户的 SecretKey，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参见 https://cloud.tencent.com/document/product/598/37140
	SecretID, SecretKey string
}

func (c *Config) New() v2.IAdapter {
	return NewCOS(c)
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
		c.CDN = c.BucketURL
	}
	return url.Parse(c.CDN)
}
