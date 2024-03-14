package fskodo

import (
	"github.com/pkg6/go-flysystem"
	"github.com/qiniu/go-sdk/v7/storage"
)

type Config struct {
	CDN        string             `json:"cdn" xml:"CDN" yaml:"CDN"`
	AccessKey  string             `json:"access_key" xml:"AccessKey" yaml:"AccessKey"`
	SecretKey  string             `json:"secret_key" xml:"SecretKey" yaml:"SecretKey"`
	Bucket     string             `json:"bucket" xml:"Bucket" yaml:"Bucket"`
	Policy     *storage.PutPolicy `json:"policy" xml:"Policy" yaml:"Policy"`
	Config     *storage.Config    `json:"config" xml:"Config" yaml:"Config"`
	PathPrefix string             `json:"path_prefix" xml:"PathPrefix" yaml:"PathPrefix"`
}

func (c *Config) New() flysystem.IAdapter {
	return NewKODO(c)
}
