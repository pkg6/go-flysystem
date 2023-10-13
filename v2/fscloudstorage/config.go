package fscloudstorage

import (
	"github.com/pkg6/go-flysystem/v2"
	"google.golang.org/api/option"
	"time"
)

var (
	DefaultWithTimeout = time.Second * 50
)

type Config struct {
	Bucket          string
	WithTimeout     time.Duration
	CredentialsFile string
	Option          []option.ClientOption
}

func (c *Config) New() v2.IAdapter {
	return NewGCS(c)
}
