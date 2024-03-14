package fscloudstorage

import (
	"github.com/pkg6/go-flysystem"
	"google.golang.org/api/option"
	"time"
)

type Config struct {
	CDN             string        `json:"cdn" xml:"CDN" yaml:"CDN"`
	Bucket          string        `json:"bucket" xml:"Bucket" yaml:"Bucket"`
	WithTimeout     time.Duration `json:"with_timeout" xml:"WithTimeout" yaml:"WithTimeout"`
	CredentialsFile string        `json:"credentials_file" xml:"CredentialsFile" yaml:"CredentialsFile"`
	CredentialsJSON string        `json:"credentials_json" xml:"CredentialsJSON" yaml:"CredentialsJSON"`
	Option          []option.ClientOption
	PathPrefix      string `json:"path_prefix" xml:"PathPrefix" yaml:"PathPrefix"`
}

func (c *Config) New() flysystem.IAdapter {
	return NewCloudStorage(c)
}
