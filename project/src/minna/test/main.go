package main

import (
	"github.com/k0kubun/pp"
	"log"
	"reflect"
)

type T struct {
	x  int
	ss []string
	m  map[string]int
}

func testDeepEqual() {
	/*
		・m2 and m3 are the same
		・the difference b/w m4 and, m2(or m3) is a value x

		note:
		・slice の場合は値と順序が同じであること
		・interface であれば実際の値が等しい否かを比較
	*/

	m1 := map[string]int{
		"a": 1,
		"b": 2,
	}

	m2 := T{
		x:  1,
		ss: []string{"one", "two"},
		m: map[string]int{
			"one": 1,
			"two": 2},
	}
	m3 := T{
		x:  1,
		ss: []string{"one", "two"},
		m: map[string]int{
			"one": 1,
			"two": 2},
	}
	m4 := T{
		x:  2,
		ss: []string{"one", "two"},
		m: map[string]int{
			"one": 1,
			"two": 2},
	}

	log.Println(reflect.DeepEqual(m1, m2))
	log.Println(reflect.DeepEqual(m2, m3))
	log.Println(reflect.DeepEqual(m3, m4))
	/*
		=>
		2021/01/01 17:50:38 false
		2021/01/01 17:50:38 true
		2021/01/01 17:50:38 false
	*/
}

func main() {
	//testDeepEqual()

	c := Client{GistGetter: &Gister{user: "RyotaBannai"}}
	if urls, err := c.ListGists(); err != nil {
		log.Println(err)
	} else {
		for _, url := range urls {
			pp.Println(url)
		}
	}
}
