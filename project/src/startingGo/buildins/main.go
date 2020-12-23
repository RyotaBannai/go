package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

var (
	constant = map[string]interface{}{
		"LOG_FILENAME":  "app_log",
		"LOG_EXTENSION": ".txt"}
	Logger *log.Logger
)

// reference: https://stackoverflow.com/questions/19965795/how-to-write-log-to-file
/*
	const (
  		Ldate         = 1 << iota  // 日付
  		Ltime                      // 時刻
  		Lmicroseconds              // 時刻のマイクロ秒
  		Llongfile                  // ソースファイル（ディレクトリパスを含む）
  		Lshortfile                 // ソースファイル（ファイル名のみ）
  		LUTC                       // タイムゾーンに依らない UTC 時刻
  		LstdFlags     = Ldate | Ltime  // 日付 (Ldata) と時刻 (Ltime) ：デフォルト
	)

	<convert interface{} to string>
	・fmt.Sprintf("%v", interface{}) // fmt
	・(interface{}).(string)         // type assertion
*/

func init() {

	var (
		today       = time.Now().Format("2000-01-01")
		logDirPath  = os.Getenv("GOPATH") + "/log/"
		logFilePath = logDirPath + today + constant["LOG_FILENAME"].(string) + constant["LOG_EXTENSION"].(string)
		file, err   = os.OpenFile(logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	)
	if err != nil {
		fmt.Println("Error: failed to create log file. \nMaybe log directory doesn't exist under GOPATH.")
		return
	}
	multiWriter := io.MultiWriter(os.Stdout, file) // 標準出力とファイル出力
	Logger = log.New(multiWriter, "Logger: ", log.LstdFlags|log.Lshortfile)
	Logger.Println("LogFile initialized.")
}

func main() {
	Logger.Println("hello world")
	getPokemon()
}
