package main

import (
	"flag"
	"fmt"
)

var (
	version string
)

func main() {
	var (
		showVersion bool
	)
	flag.BoolVar(&showVersion, "v", false, "show version")
	flag.BoolVar(&showVersion, "version", false, "show version")
	flag.Parse() // 引数からオプションをパース
	if showVersion {
		fmt.Println("version: ", version)
		return
	}
	// some application running...
}
