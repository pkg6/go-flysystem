package fsbos

import "github.com/pkg6/go-flysystem"

type Config struct {
	CDN              string
	Ak               string
	Sk               string
	Endpoint         string
	RedirectDisabled bool
	Bucket           string
	PathPrefix       string
}

func (c *Config) New() flysystem.IAdapter {
	return NewBOS(c)
}
