package main

import (
	"time"
	"fmt"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(i)
	}
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func run() {
	go say("world") // 先頭に go とつけるだけで立ち上げる軽量スレッドが立ち上がる
	say("hello")
}

func dist_tasks() {
	// ２つの goroutine 間で作業を分配する!!
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c
	fmt.Println(x, y, x+y)
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func closeChan() {
	c := make(chan int, 10) // channel をバッファとして使用. ２つ目の引数のバッファの長さを与える.
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

func main() {
	closeChan()
}
