package main

import (
	"bytes"
	"fmt"
	"log"
)

func Logger() (buf bytes.Buffer, infof func(message string)) {
	var (
		logger = log.New(&buf, "INFO: ", log.Lshortfile)
	)
	infof = func(info string) {
		logger.Output(2, info)
	}
	return
}

func main() {
	lbuf, logger := Logger()
	logger("hello world")
	fmt.Println(lbuf)

	log.Print("Hello, world!")

	checkFileInfo()
}
