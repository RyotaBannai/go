package main

import (
	"log"
	"os/exec"
	"runtime"
)

func runScript() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "myapp.bat")
	} else {
		cmd = exec.Command("bin/sh", "-c", "myapp.sh")
	}
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
