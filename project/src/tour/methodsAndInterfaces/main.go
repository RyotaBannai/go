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
func (v *Vertex) Abs() float64 { // distance
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

// interface は'メソッドシグネチャ'の集まりで定義される
type Abser interface {
	Abs() float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	} else {
		return float64(f)
	}
}

func letImplementInteraface() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}
	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	//a = v // Error! Abs メソッドは、Vertex ではなく、*Vertex の定義であるため.
	fmt.Println(a.Abs())
}

func main() {
	myMethod()
}
