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

/*
	https://stackoverflow.com/questions/34931059/go-tutorial-select-statement
	ポイントは unbuffered がどのようになれば、データが流れるかを理解すること.
*/
func fibonacci2(c, quit chan int) {
	x, y := 0, 1
	for { // this is equivalent to a while loop, without a stop condition
		select {
		case c <- x: // when we can send to channel c, and because c is unbuffered, we can only send to channel c when someone tries to receive from it
			x, y = y, x+y
		case <-quit: // when we can receive from channel quit, and because quit is unbuffered, we can only receive from channel quit when someone tries to send to it
			fmt.Println("quit")
			return
		}
	}
}

func mySelect() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	// fibonacci2 は select を宣言して channel 振り分けをしているだけ
	fibonacci2(c, quit)
}

func main() {
	mySelect()
}
