package main

import (
	"fmt"
	"math"
)

func main() {
	a := []string{
		"Michael",
		"Jobs",
		"Mark"} // 最終行は } か、カンマを入れないとコンパイラが自動でセミコロンを入れるためエラー
	fmt.Println(a)
	fmt.Println(pkgVar) // print package variable // go にはグローバル変数はない
	checkedOverflow(1, math.MaxUint32)
}
