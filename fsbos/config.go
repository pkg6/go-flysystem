package fsbos

import "github.com/pkg6/go-flysystem"

type Config struct {
	CDN string `json:"cdn" xml:"CDN" yaml:"CDN"`
	Ak  string `json:"ak" xml:"Ak" yaml:"Ak"`
	Sk  string `json:"sk" xml:"Sk" yaml:"Sk"`
	//https://cloud.baidu.com/doc/BOS/s/Ojwvyrpgd
	Endpoint         string `json:"endpoint" xml:"Endpoint" yaml:"Endpoint"`
	RedirectDisabled bool   `json:"redirect_disabled" xml:"RedirectDisabled" yaml:"RedirectDisabled"`
	Bucket           string `json:"bucket" xml:"Bucket" yaml:"Bucket"`
	PathPrefix       string `json:"path_prefix" xml:"PathPrefix" yaml:"PathPrefix"`
}

func (c *Config) New() flysystem.IAdapter {
	return NewBOS(c)
}
