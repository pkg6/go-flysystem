package fskodo

import (
	"github.com/pkg6/go-flysystem"
	"github.com/qiniu/go-sdk/v7/storage"
)

type Config struct {
	CDN                  string
	AccessKey, SecretKey string
	Bucket               string
	Policy               *storage.PutPolicy
	Config               *storage.Config
	PathPrefix           string
}

func (c *Config) New() flysystem.IAdapter {
	return NewKODO(c)
}
