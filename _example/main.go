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
