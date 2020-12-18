package main

import "fmt"

func main() {
	a := []string{
		"Michael",
		"Jobs",
		"Mark"} // 最終行は } か、カンマを入れないとコンパイラが自動でセミコロンを入れるためエラー
	fmt.Println(a)
}
