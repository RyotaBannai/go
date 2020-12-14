package main

import "fmt"

type I1 interface {
	M()
}

type T1 struct {
	S string
}

func (t *T1) M() {
	if t == nil {
		fmt.Println("<nil>")
	} else {
		fmt.Println(t.S)
	}
}

func main() {
	var i I1
	var t *T1 // T1 does not implement I1, so can't be assigned to i
	i = t
	describe(i)

	i = &T1{"hello"} //
	describe(i)
	i.M()

}
func describe(i I1) {
	fmt.Printf("(%v, %T)\n", i, i)
}
