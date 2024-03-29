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

## example

~~~
package main

import (
	"fmt"
	"github.com/pkg6/go-flysystem"
	"github.com/pkg6/go-flysystem/local"
	"strings"
)

func main() {
	//Define the root directory of the local adapter
	root := "./_example/test_data"
	// Create local adapter
	localAdapter := local.New(&local.Config{Root: root})
	//Initialize the adapter
	adapters := flysystem.NewAdapters(localAdapter)
	adapters.Extend(local.New(&local.Config{Root: "./_example/test_data/2"}), "local2")
	var err error
	_, err = adapters.WriteReader("4.txt", strings.NewReader("test"))
	fmt.Println(err)
	adapter, err := adapters.Adapter("local2")
	_, err = adapter.WriteReader("4.txt", strings.NewReader("test"))
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
~~~

## Associated Projects

gfs： [github.com/pkg6/gfs](https://github.com/pkg6/gfs)
