package main

import "fmt"

// type assertions
func typeAssert() {
	var i interface{} = "hello"
	// インターフェースの値 i が具体的な型（string）を保持し基になる string の値を変数 s に代入することを主張している.
	// この場合、i が string 型を保持していない場合、この分は panic を起こす.
	s := i.(string)
	fmt.Println(s)

	// インターフェースの値が特定の型を保持しているかテスト
	ss, ok := i.(string)
	if ok {
		fmt.Println(ss)
	} else {
		fmt.Println("actually it's a panic on this operation...")
	}
}

func main() {
	typeAssert()
}
