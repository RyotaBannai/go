package main

import (
	"strings"
	"fmt"
	"io"
)

func main() {
	r := strings.NewReader("Hello, World!")
	b := make([]byte, 8)
	for {
		n, err := r.Read(b) // 先頭から 8 bytes 取り出す
		fmt.Printf("n = %v, err = %v, b = %b\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n]) // %q -> double quotes
		if err == io.EOF {
			break
		}
	}
}
