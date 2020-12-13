package main

import (
	"math"
	"fmt"
)

type Vertex struct {
	X, Y float64
}

// methods -> レシーバ引数を伴う関数 methods != functions
// Vertex 型の receiver を持つ Abs method
func (v Vertex) Abs() float64 { // distance
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Pointer receivers
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func myMethod() {
	v := Vertex{3, 4}    // 5
	v.Scale(10)          // 自分自身を更新
	fmt.Println(v.Abs()) // 50
}

func main() {
	myMethod()
}
