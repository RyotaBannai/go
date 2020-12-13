package main

import (
	"fmt"
)

type BasicStruct struct {
	i int
	b bool
}

type Vertex struct {
	X int
	Y int
}

type LatLong struct {
	Lat, Long float64
}

func main() {
	myMap()
}

func myMap() {
	m := make(map[string]LatLong)
	m["Bell Labs"] = LatLong{
		40.58433, -74.39967,
	}
	fmt.Println(m)

	// その型がリテラルの要素から推測できるため、省略することができる.
	m2 := map[string]LatLong{
		"Google": {37.42202, -122.08408},
	}
	fmt.Println(m2)

}

func myIterate() {
	s := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	// もしインデックスだけが必要なのであれば、2つ目の値を省略 for i := range s{...}
	for i, v := range s {
		fmt.Printf("index: %d, value: %v \n", i, v)
	}
}

//
//var warehouse [][]uint8
//
//func Pic(dx, dy int) [][]uint8 {
//	append(warehouse, uint8([dx, dy]))
//	return warehouse
//}

func myAppend() {
	// 変数群を追加する際に元の slice の容量が小さい場合は、より大きいサイズの配列を割り当て直す
	var s []int
	s = append(s, 1, 2, 3)
	fmt.Println(s)

}

func myArrays() {
	// slice は元の配列への参照であり、どのデータも格納していない -> slice の要素を更新すると元の配列も更新される.
	primes := [6]int{2, 3, 5, 7, 11, 13}
	mySlice := primes[1:3] // [3,5] -> Python みたいに下限上限を省略することもできる.
	fmt.Println(mySlice)

	structSlice := []BasicStruct{
		{1, true},
		{2, false},
		{3, true},
		{4, false},
	}
	fmt.Println(structSlice)
	printSlice(structSlice)

	// slice のゼロ値は nil
	var s []int
	if s == nil {
		fmt.Println("ゼロ値は nil!")
	}
}

func sliceViaMake() {
	a := make([]int, 5) // len(a)=5, cap(a)=5
	fmt.Println(cap(a))

	// make の3番目の引数(third operand)に、スライスの容量(capacity)を指定できる
	b := make([]int, 0, 5) // len(b)=0, cap(b)=5
	b = b[:cap(b)]         // len(b)=5, cap(b)=5
	b = b[1:]              // len(b)=4, cap(b)=4
}

func printSlice(s []BasicStruct) {
	// format の %v は value を表示できる.
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
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
