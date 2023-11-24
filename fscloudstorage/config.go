package fscloudstorage

import (
	"github.com/pkg6/go-flysystem"
	"google.golang.org/api/option"
	"time"
)

type Config struct {
	CDN             string
	Bucket          string
	WithTimeout     time.Duration
	CredentialsFile string
	Option          []option.ClientOption
	PathPrefix      string
}

func (c *Config) New() flysystem.IAdapter {
	return NewCloudStorage(c)
}
