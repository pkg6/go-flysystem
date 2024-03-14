package fscos

import "github.com/pkg6/go-flysystem"

type Config struct {
	CDN string `json:"cdn" xml:"CDN" yaml:"CDN"`
	// 存储桶名称，由 bucketname-appid 组成，appid 必须填入，可以在 COS 控制台查看存储桶名称。 https://console.cloud.tencent.com/cos5/bucket
	// 替换为用户的 region，存储桶 region 可以在 COS 控制台“存储桶概览”查看 https://console.cloud.tencent.com/
	BucketURL string `json:"bucket_url" xml:"BucketURL" yaml:"BucketURL"`
	// 用户的 SecretId，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参见 https://cloud.tencent.com/document/product/598/37140
	// 用户的 SecretKey，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参见 https://cloud.tencent.com/document/product/598/37140
	SecretID   string `json:"secret_id" xml:"SecretID" yaml:"SecretID"`
	SecretKey  string `json:"secret_key" xml:"SecretKey" yaml:"SecretKey"`
	PathPrefix string `json:"path_prefix" xml:"PathPrefix" yaml:"PathPrefix"`
}

func (c *Config) New() flysystem.IAdapter {
	return NewCOS(c)
}
