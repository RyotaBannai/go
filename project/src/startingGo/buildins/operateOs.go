package main

import (
	"fmt"
	"github.com/TylerBrock/colorjson"
	"io/ioutil"
	"os"
)

func prettyPrinter(data map[string]interface{}) {
	f := colorjson.NewFormatter()
	f.Indent = 2
	s, _ := f.Marshal(data)
	fmt.Println(string(s))
}

func fileCloser(file *os.File) {
	if err := file.Close(); err != nil {
		fmt.Println("error occurred while closing file.")
	} else {
		fmt.Println("closed file successfully!")
	}
}

func fileOpener(filepath string) (file *os.File, err error) {
	if file, err = os.Open(filepath); err != nil { // file に assign するのは メモリ上の実態データ
		fmt.Println("file doesn't exist.")
		return nil, err
	} else {
		fmt.Println("file exists.")
		return
	}
}

func checkFileInfo() {
	file, err := fileOpener("./sentence.txt")
	defer fileCloser(file) // return の前に宣言
	if err != nil {
		return
	}

	// bs := make([]byte, 128)
	// // n は実際に読み込んだバイト数
	// if n, err := file.Read(bs); err != nil {
	// 	fmt.Println("error")
	// } else {
	// 	fmt.Println(n)
	// }

	// file のステータスを確認

	if fi, err := file.Stat(); err != nil {
	} else {
		fmt.Println(fi.Name())    // ファイル名
		fmt.Println(fi.Size())    // ファイルサイズ int64
		fmt.Println(fi.Mode())    // ファイルのモード os.FileMode (chmod)
		fmt.Println(fi.ModTime()) // ファイルの最終更新時間 time.Time
		fmt.Println(fi.IsDir())   // ディレクトリかどうか
	}

	if b, err := ioutil.ReadAll(file); err != nil {
	} else {
		fmt.Println(string(b))
	}
}