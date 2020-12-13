package main

import "fmt"

func main() {
	f := fibonatti()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

// fibonacci?
func fibonatti() func() int {
	first := -1
	second := 1
	var temp int
	return func() int {
		temp = first
		first = second
		second = temp + first
		return second
	}
}
