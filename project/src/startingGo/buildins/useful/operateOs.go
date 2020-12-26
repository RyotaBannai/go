package useful

import (
	"bufio"
	"fmt"
	"github.com/TylerBrock/colorjson"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

/*
	read file reference: http://zetcode.com/golang/readfile/
*/

func PrettyPrinter(data map[string]interface{}) {
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

func checkNumberOfBytes(file *os.File) int {
	if n, err := file.Read(make([]byte, 128)); err != nil { // n は実際に読み込んだバイト数
		fmt.Println("error")
		return 0
	} else {
		fmt.Println(n)
		return n
	}
}

func printFileStats(file *os.File) {
	// file のステータス
	if fi, err := file.Stat(); err != nil {
	} else {
		fmt.Println(fi.Name())    // ファイル名
		fmt.Println(fi.Size())    // ファイルサイズ int64
		fmt.Println(fi.Mode())    // ファイルのモード os.FileMode (chmod)
		fmt.Println(fi.ModTime()) // ファイルの最終更新時間 time.Time
		fmt.Println(fi.IsDir())   // ディレクトリかどうか
	}
}

func readContents(file *os.File) string {
	if b, err := ioutil.ReadAll(file); err != nil {
		return ""
	} else {
		return string(b)
	}
}

func readContents2(filePath string) (string, error) {
	// file を close する必要ない.
	if content, err := ioutil.ReadFile(filePath); err != nil {
		fmt.Println("Failed to read content.")
		return "", err
	} else {
		return string(content), nil
	}
}

func printContentsByScanner(file *os.File) {
	// read content by newline
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(strings.Repeat("-", 5))
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func WordCounter(file *os.File) map[string]int {
	var (
		wordMap = make(map[string]int)
		scanner = bufio.NewScanner(file)
	)
	scanner.Split(bufio.ScanWords) // 引数は splitFunc
	for scanner.Scan() {
		word := scanner.Text()
		if _, ok := wordMap[word]; ok {
			wordMap[word] += 1
		} else {
			wordMap[word] = 1
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return nil
	}
	return wordMap
}

func readDirRecursively(dir *os.File) {
	if files, err := dir.Readdir(0); err != nil { // 0 で全て取得
	} else {
		for _, file := range files {
			if file.IsDir() {
				fmt.Println(file.Name()) // ディレクトリのみ取得
			}
		}
	}
}

func checkFile() {
	file, err := fileOpener("./sentence.txt")
	defer fileCloser(file) // return の前に宣言
	if err != nil {
		return
	}
	counter := WordCounter(file)
	fmt.Println(counter)

	if content, err := readContents2("./sentence.txt"); err != nil {
	} else {
		fmt.Println(content)
	}
}

func checkDir() {
	dir, err := fileOpener("../") // startingGo
	defer fileCloser(dir)
	if err != nil {
	} else {
		readDirRecursively(dir)
	}
}
