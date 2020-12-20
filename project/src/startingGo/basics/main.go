package main

import (
	"encoding/json"
	"fmt"
	"github.com/TylerBrock/colorjson"
	"io/ioutil"
	"math"
	"os"
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

func caseByType(x interface{}) {
	switch v := x.(type) {
	case int, uint:
		fmt.Printf("%T, %d\n", v, v) // v * v はできない → 複数列挙した場合、型判定できないため、interface{} 型となってしまうため.
	default:
		fmt.Println("Unknown type.")
	}
}

func runDefer() {
	/*
		・defer に登録された式は関数の終了後に評価される.
		・複数 defer を宣言した場合、初めに defer 宣言された関数は最後に評価されることに注意.
	*/

	//defer fmt.Println("Deferred.") // ok as well
	defer func() {
		fmt.Println("Deferred.")
	}()
	fmt.Println("Done.")
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func prettyPrint(data map[string]interface{}) {
	f := colorjson.NewFormatter()
	f.Indent = 2
	s, _ := f.Marshal(data)
	fmt.Println(string(s))
}

func openFile() {
	jsonFile, err := os.Open("./users.json")
	if err != nil {
		fmt.Println("file doesn't exist.")
		return
	} else {
		fmt.Println("file exists.")
	}
	// check(err)

	// structure json data.
	// https://tutorialedge.net/golang/parsing-json-with-golang/
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var result map[string]interface{}
	unmarshalErr := json.Unmarshal([]byte(byteValue), &result)
	if unmarshalErr != nil {
		fmt.Println("unmarshal Err occurred.")
	} else {
		prettyPrint(result)
	}

	defer func() {
		err := jsonFile.Close()
		if err != nil {
			fmt.Println("error occurred while closing file..")
		} else {
			fmt.Println("closed file successfully")
		}
	}()
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
	caseByType(uint(10))
	runDefer()
	openFile()
}
