package config

import (
	"github.com/pkg6/go-flysystem/fsbos"
	"github.com/pkg6/go-flysystem/fscloudstorage"
	"github.com/pkg6/go-flysystem/fscos"
	"github.com/pkg6/go-flysystem/fskodo"
	"github.com/pkg6/go-flysystem/fsoss"
	"github.com/pkg6/go-flysystem/local"
)

type Config struct {
	LOCAL        *local.Config          `gfs:"local,default"`
	OSS          *fsoss.Config          `gfs:"oss"`
	BOS          *fsbos.Config          `gfs:"bos"`
	COS          *fscos.Config          `gfs:"cos"`
	KODO         *fskodo.Config         `gfs:"kodo"`
	CloudStorage *fscloudstorage.Config `gfs:"cloud_storage"`
}
