package flysystem_test

import (
	"github.com/pkg6/go-flysystem"
	"github.com/pkg6/go-flysystem/config"
	"github.com/pkg6/go-flysystem/fsbos"
	"github.com/pkg6/go-flysystem/fscloudstorage"
	"github.com/pkg6/go-flysystem/fscos"
	"github.com/pkg6/go-flysystem/fskodo"
	"github.com/pkg6/go-flysystem/fsoss"
	"github.com/pkg6/go-flysystem/local"
	"github.com/pkg6/gfs/bosfs"
	"google.golang.org/api/option"
	"testing"
)

func TestNewConfig(t *testing.T) {
	cfg := config.Config{
		LOCAL: &local.Config{Root: "./_example/test_data"},
		OSS:   &fsoss.Config{},
		BOS: &fsbos.Config{
			Endpoint: bosfs.DefaultEndpoint,
			Ak:       "Ak",
			Sk:       "Sk",
			Bucket:   "test bucket",
		},
		COS: &fscos.Config{
			BucketURL: "https://bucket-id.cos.ap-beijing.myqcloud.com",
			SecretID:  "SecretID",
			SecretKey: "SecretKey",
		},
		KODO: &fskodo.Config{
			AccessKey: "AccessKey",
			SecretKey: "SecretKey",
			Bucket:    "test bucket",
		},
		CloudStorage: &fscloudstorage.Config{
			Bucket: "test bucket",
			Option: []option.ClientOption{
				option.WithCredentialsFile("CredentialsFile.json"),
			},
		},
	}
	fs, err := flysystem.NewConfig(&cfg)
	if err != nil {
		t.Fatal(err)
	}
	if fs.Disk("") != "local" {
		t.Fatal(err)
	}
	if fs.Disk("oss") != "oss" {
		t.Fatal(err)
	}
	if fs.Disk("oss2") == "" {
		t.Fatal(err)
	}
	if fs.DiskExist("oss2") == true {
		t.Fatal(err)
	}
}
