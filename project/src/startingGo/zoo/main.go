package main

/*
	・同じディレクトリには package 名は一つまでであるが、その名前はディレクトリ名と一致していなくても良い.
	・import に指定するパッケージは通常 GOPATH に指定されたディレクトリ内のパッケージから探索される.
*/

import (
	"fmt"
	"starginGo/zoo/animals" // 相対パスだと package ビルドできない...
)

func main() {
	fmt.Println(animals.Elephant()) // Grass
	fmt.Println(AppName())
}
