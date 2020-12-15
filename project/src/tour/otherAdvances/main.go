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

// type switches

func typeSwitch(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	typeAssert()
	typeSwitch(21)
	typeSwitch("hello")
	typeSwitch(true)
}
