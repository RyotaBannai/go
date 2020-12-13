package main

import (
	"fmt"
)

func main() {
	myStructs()
}

// js とか Object みたいなものか.
type Vertex struct {
	X int
	Y int
}

func myStructs() {
	v := Vertex{1, 2}
	v.X = 0
	fmt.Println(v)

	// Pointers to structs
	p := &v        // using pointer
	p.X = 1e9      // (*p).X のように記述もできる...
	fmt.Println(v) // 更新される.

	// Name: 構文を使って、structs の一部だけを列挙することができる
	p2 := Vertex{X: 10} // 初期化しないときは int type の場合は 0 になる.
	fmt.Println(p2)
}

func basics() {
	num := 123
	str := "ABD"
	fmt.Println("Hello, World")
	fmt.Printf("num=%d, str=%s", num, str)
	var huge complex64 = 10 + 10i
	fmt.Println(huge)
}
