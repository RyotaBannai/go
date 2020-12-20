package main

import (
	"fmt"
	"math"
)

func loopString() {
	for i, r := range "あいうえお" {
		fmt.Printf("[%d] = %d\n", i, r)
	}
	/*
		文字列に対する range は ulf-8 でエンコードされた文字列のコードポイントごとに反復される.
		→ 第一変数の値はインデックスではなく、コードポインタが開始されるバイト列のインデックス.
		→ 文字のコードポイントに応じて文字列のインデックスの増分は異なる.
		[0] = 12354
		[3] = 12356
		[6] = 12358
		[9] = 12360
		[12] = 12362
	*/
}

func main() {
	a := []string{
		"Michael",
		"Jobs",
		"Mark"} // 最終行は } か、カンマを入れないとコンパイラが自動でセミコロンを入れるためエラー
	fmt.Println(a)
	fmt.Println(pkgVar) // print package variable // go にはグローバル変数はない
	checkedOverflow(1, math.MaxUint32)
	loopString()
}
