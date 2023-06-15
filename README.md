# go-flysystem

[![Go Report Card](https://goreportcard.com/badge/github.com/pkg6/go-flysystem)](https://goreportcard.com/report/github.com/pkg6/go-flysystem)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/github.com/pkg6/go-flysystem?tab=doc)
[![Sourcegraph](https://sourcegraph.com/github.com/pkg6/go-flysystem/-/badge.svg)](https://sourcegraph.com/github.com/pkg6/go-flysystem?badge)
[![Release](https://img.shields.io/github/release/pkg6/go-flysystem.svg?style=flat-square)](https://github.com/pkg6/go-flysystem/releases)
[![Goproxy.cn](https://goproxy.cn/stats/github.com/pkg6/go-flysystem/badges/download-count.svg)](https://goproxy.cn)


## About Flysystem

Flysystem is a file storage library for Golang. It provides one interface to interact with many types of filesystems. When you use Flysystem, you're not only protected from vendor lock-in, you'll also have a consistent experience for which ever storage is right for you.

## Install

~~~
$ go get github.com/pkg6/go-flysystem
~~~

## Officially supported adapters

- **[Local](https://github.com/pkg6/go-flysystem/blob/main/local/)**
- **[OSS](https://help.aliyun.com/product/31815.html)**

## example

~~~
package main

import (
	"fmt"
	"github.com/pkg6/go-flysystem"
	"github.com/pkg6/go-flysystem/fsoss"
	"github.com/pkg6/go-flysystem/local"
	"strings"
)

func main() {
	//Define the root directory of the local adapter
	root := "./_example/test_data"
	// Create local adapter
	localAdapter := local.New(root)
	ossAdapter := fsoss.New(fsoss.Config{
		Bucket:          "test",
		Endpoint:        "oss-cn-hangzhou.aliyuncs.com",
		AccessKeyID:     "*******************",
		AccessKeySecret: "**************",
		PathPrefix:      "shop",
	})
	//Initialize the adapter
	adapters := flysystem.NewAdapters(localAdapter)
	adapters.Extend(ossAdapter.DiskName(), ossAdapter)
	var err error
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
	//Create directory
	err = adapters.CreateDirectory("test1/test12")
	fmt.Println(err)
	//Delete directory
	_, err = adapters.DeleteDirectory("test1/test12")
	fmt.Println(err)
}
~~~

> If Disk is not specified, it will be executed using the first registered driver



You can always [create an adapter](https://github.com/pkg6/go-flysystem/blob/main/adapter.go) yourself.

~~~
type IFS interface {
	// Exists Determine if the file exists
	Exists(path string) (bool, error)
	// Size Get File Size
	Size(path string) (int64, error)
	// Write write file content and return full path
	Write(path string, contents []byte) (string, error)
	// WriteStream Resource file write returns full path
	WriteStream(path, resource string) (string, error)
	// Update Update the file content and return the updated full path
	Update(path string, contents []byte) (string, error)
	// UpdateStream Return the updated full path based on resource file updates
	UpdateStream(path, resource string) (string, error)
	// Read Read file
	Read(path string) ([]byte, error)
	// Delete  Deleting files returns the number of deleted files
	Delete(path string) (int64, error)
	// DeleteDirectory Number of files deleted from the deleted directory
	DeleteDirectory(dirname string) (int64, error)
	// CreateDirectory create directory
	CreateDirectory(dirname string) error
	// MimeType Get File MimeType
	MimeType(path string) (string, error)
	// Move move file
	Move(source, destination string) (bool, error)
	// Copy copy file
	Copy(source, destination string) (bool, error)
}
type IAdapter interface {
	IFS
	// DiskName Default Disk Name
	DiskName() string
	// Clone Initialization parameters
	Clone() IAdapter
}
~~~

