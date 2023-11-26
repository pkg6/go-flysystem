package main

import (
	"fmt"
	"github.com/pkg6/go-flysystem"
	"github.com/pkg6/go-flysystem/config"
	"github.com/pkg6/go-flysystem/fsbos"
	"github.com/pkg6/go-flysystem/fscloudstorage"
	"github.com/pkg6/go-flysystem/fscos"
	"github.com/pkg6/go-flysystem/fskodo"
	"github.com/pkg6/go-flysystem/fsoss"
	"github.com/pkg6/go-flysystem/local"
	"github.com/zzqqw/gfs/bosfs"
	"google.golang.org/api/option"
	"strings"
)

func main() {
	c := config.Config{
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
	//Define the root directory of the local adapter
	root := "./_example/test_data"
	//Initialize the adapter
	adapters, err := flysystem.NewConfig(&c)
	fmt.Println(err)
	_, err = adapters.WriteReader("4.txt", strings.NewReader("test"))
	fmt.Println(err)
	//Write file
	_, err = adapters.Write("1.txt", []byte("test data"))
	fmt.Println(err)
	//Write data from resource file
	_, err = adapters.WriteStream("2.txt", root+"/1.txt")
	fmt.Println(err)
	//Update file
	_, err = adapters.Update("1.txt", []byte("test update data"))
	fmt.Println(err)
	//Update data from resource file
	_, err = adapters.UpdateStream("2.txt", root+"/1.txt")
	fmt.Println(err)
	exists, _ := adapters.Exists("2.txt")
	fmt.Println(exists)
	//Read file
	read, err := adapters.Read("2.txt")
	fmt.Println(read, err)
	//Get file mime type
	mimeType, err := adapters.MimeType("2.txt")
	fmt.Println(mimeType, err)
	//Get file size
	size, err := adapters.Size("2.txt")
	fmt.Println(size, err)
	//Move file
	_, err = adapters.Move("1.txt", "4.txt")
	fmt.Println(err)
	//Copy file
	_, err = adapters.Copy("2.txt", "5.txt")
	fmt.Println(err)
}
