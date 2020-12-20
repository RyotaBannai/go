package main

import (
	"fmt"
	"runtime"
)

func showRuntimeDetails() {
	fmt.Printf("NumCPU: %d\n", runtime.NumCPU())             // 4
	fmt.Printf("NumGoroutine: %d\n", runtime.NumGoroutine()) // 1
	fmt.Printf("Version: %s\n", runtime.Version())           // go1.15.6
}
