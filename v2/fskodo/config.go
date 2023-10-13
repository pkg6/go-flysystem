package fskodo

import (
	"github.com/pkg6/go-flysystem/v2"
	"github.com/qiniu/go-sdk/v7/storage"
)

type Config struct {
	AccessKey, SecretKey string
	Bucket               string

	Policy *storage.PutPolicy
	Config *storage.Config
}

func (c *Config) New() v2.IAdapter {
	return NewKoDo(c)
}
