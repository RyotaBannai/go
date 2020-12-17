package main

import (
	"io"
	"strings"
	"os"
)

type rot13Reader struct {
	r io.Reader
}

func (a *rot13Reader) Read(rb []byte) (n int, e error) {
	n, e = a.r.Read(rb)
	if e == nil {
		for i, v := range rb {
			// 初めに a A で差分をとってベースを 0 にする.
			// %26 することで 26 を超えた byte を上段の 13 に変換.
			switch {
			case v >= 'A' && v <= 'Z':
				rb[i] = (v-'A'+13)%26 + 'A'
			case v >= 'a' && v <= 'z':
				rb[i] = (v-'a'+13)%26 + 'a'
			}
		}
	}
	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
