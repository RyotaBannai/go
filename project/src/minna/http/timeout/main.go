package main

import (
	"bytes"
	"context"
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

func httpGetWithContext(url, method string, dst io.Writer) error {
	var (
		client      = &http.Client{}
		ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
		req, _      = http.NewRequest(method, url, nil)
		resp, err   = client.Do(req.WithContext(ctx))
	)
	if err != nil {

		// Get "http://localhost:8080/slow": context deadline exceeded
		return err
	}
	defer cancel()
	defer resp.Body.Close()

	// Get "http://localhost:8080/slow": context deadline exceeded
	_, err = io.Copy(dst, resp.Body)
	return err
}

func main() {
	err := httpGetWithContext("http://localhost:8080/slow", "GET", os.Stdout)
	if err != nil {
		io.Copy(os.Stderr, bytes.NewBufferString(fmt.Sprintf("%v\n", err)))
		/*
			NewBufferString: creates and initializes a new Buffer using string s as its initial contents.
			It is intended to prepare a buffer to read an existing string.
		*/
	}
}
