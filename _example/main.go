package main

import (
	"fmt"
	"github.com/pkg6/go-flysystem"
	"github.com/pkg6/go-flysystem/aliyunoss"
	"github.com/pkg6/go-flysystem/googlecloudstorage"
	"github.com/pkg6/go-flysystem/local"
	"google.golang.org/api/option"
	"strings"
)

var (
	root          = "./_example/test_data"
	localAdapter  flysystem.IAdapter
	local2Adapter flysystem.IAdapter
	ossAdapter    flysystem.IAdapter
	google        flysystem.IAdapter
)

func init() {
	localAdapter = local.New(root)
	local2Adapter = local.New("./_example/test_data/2")
	ossAdapter = aliyunoss.New(&aliyunoss.Config{
		Bucket:          "test",
		Endpoint:        "oss-cn-hangzhou.aliyuncs.com",
		AccessKeyID:     "*******************",
		AccessKeySecret: "**************",
		PathPrefix:      "test",
	})
	google = googlecloudstorage.New(&googlecloudstorage.Config{
		Bucket: "test bucket",
		Option: []option.ClientOption{
			option.WithCredentialsFile("CredentialsFile.json"),
		},
	})
}

func main() {
	adapters := flysystem.NewAdapters()
	adapters.Extend(localAdapter)
	adapters.Extend(ossAdapter)
	adapters.Extend(local2Adapter, "local2")
	adapters.Extend(google)
	var err error
	_, err = adapters.WriteReader("4.txt", strings.NewReader("test"))
	fmt.Println(err)
	_, err = adapters.Disk("local2").WriteReader("4.txt", strings.NewReader("test"))
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
	if err != nil {
		return
	}
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
