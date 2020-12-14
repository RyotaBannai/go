package main

import "fmt"

type I interface {
	// interface は method シグネチャの集まり
	M()
}

type T struct {
	S string
}

// これは構造体 T が暗黙的に interface I を実装していることを表す.
// (つまり、特に他言語のように implement のような宣言をする必要はない.s)
func (t T) M() {
	fmt.Println(t.S)
}

func call() {
	var i I = T{"Hello"} // 構造体に型宣言は必要なけど、interface には必要 like i I = ..
	i.M()
}

func main() {
	call()
}
