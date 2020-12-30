package main

import (
	"fmt"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fmt.Fprintf(w, "GET: hello, world\n")
	case http.MethodPost:
		fmt.Fprintf(w, "POST: hello, world\n")
	case http.MethodPut:
		fmt.Fprintf(w, "PUT: hello, world\n")
	case http.MethodDelete:
		fmt.Fprintf(w, "DELETE: hello, world\n")
	default:
		fmt.Fprintf(w, "DEFAULT: hello, world\n")
	}
}

func slowHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(10 * time.Second)
	handler(w, r)
}

func main() {
	http.HandleFunc("/", handler)     // ハンドラを登録
	http.HandleFunc("/slow", slowHandler)
	http.ListenAndServe(":8080", nil)
}
