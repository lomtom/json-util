package main

import (
	"github.com/tidwall/gjson"

	"flag"
	"fmt"
	"os"
)

var file = flag.String("f", "", "file path")
var key = flag.String("k", "", "key name")

func main() {
	load()
	bytes, err := os.ReadFile(*file)
	if err != nil {
		fmt.Println("load file error,", err)
		os.Exit(1)
	}
	value := gjson.GetBytes(bytes, *key)
	fmt.Println(value.String())
}

func load() {
	flag.Parse()
	if file == nil || *file == "" {
		fmt.Println("file is required")
		os.Exit(1)
	}
	if key == nil || *key == "" {
		fmt.Println("key is required")
		os.Exit(1)
	}
}
