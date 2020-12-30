package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func httpGet(url, method string, dst io.Writer) error {
	client := &http.Client{
		Timeout: 5 * time.Second}
	req, _ := http.NewRequest(method, url, nil)
	resp, err := client.Do(req)
	if err != nil {

		// レスポンスヘッダ取得まで 5 秒経過したらエラー
		// Get "http://localhost:8080/slow": context deadline exceeded (Client.Timeout exceeded while awaiting headers)
		return err
	}

	defer resp.Body.Close()

	// レスポンスヘッダ取得まで 5 秒経過したらエラー
	// Get "http://localhost:8080/slow": context deadline exceeded (Client.Timeout exceeded while awaiting headers)
	_, err = io.Copy(dst, resp.Body)
	return err
}

func main() {
	err := httpGet("http://localhost:8080/slow", "GET", os.Stdout)
	if err != nil {
		fmt.Println(err)
	}
}
