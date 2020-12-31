package main

import (
	"fmt"
	"minna/reflect/other"
	"reflect"
)

func main() {
	p := other.Me
	rt := reflect.TypeOf(p)
	rv := reflect.ValueOf(p)

	/*
		private な field には reflect でもアクセスできない
	*/

	// field を一つずつ取り出す
	// methods も同様に取り出すことができる
	for i := 0; i < rv.NumField(); i++ {
		ft := rt.Field(i)
		fv := rv.Field(i)
		/*
			export されていれば PkgPath は空になる
		*/
		if path := ft.PkgPath; path == "" {
			fmt.Printf("ft(%d) -> %#v\n", i, ft) // reflect.StructField...
			fmt.Printf("ft(%d) -> %#v\n", i, fv.Interface())
		} else {
			fmt.Println(path)
		}
	}
}
