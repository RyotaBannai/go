package main

import (
	"context"
	"fmt"
	"sync"
)

/*
	・goroutine を終了させる方法として、他には close もある

	for url, more := range <- queue{
		if more {}
		else{
			wg.Done()
			return
		}
	}
*/

var (
	wg = &sync.WaitGroup{}
)

func cancelableGoroutine() {
	ctx, cancel := context.WithCancel(context.Background())
	queue := make(chan string)
	for i := range N(2) {
		wg.Add(1)
		go worker(ctx, queue, i)
	}
	for _, url := range []string{
		"https://www.example.com",
		"https://www.example.net",
		"https://www.example.com/foo",
		"https://www.example.com/bar"} {
		queue <- url
	}
	cancel()
	wg.Wait()
}

func worker(ctx context.Context, queue chan string, workerId int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%v: Context Done.\n", workerId)
			wg.Done()
			return // do not forget unless goroutine won't stop.
		case value := <-queue:
			fmt.Printf("%v: Get value from queue: %s\n", workerId, value)
		}
	}
}
