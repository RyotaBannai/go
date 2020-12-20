package main

import (
	"fmt"
	"runtime"
	"time"
)

func showRuntimeDetails() {
	fmt.Printf("NumCPU: %d\n", runtime.NumCPU())             // 4
	fmt.Printf("NumGoroutine: %d\n", runtime.NumGoroutine()) // 1
	fmt.Printf("Version: %s\n", runtime.Version())           // go1.15.6
}

func receiver(name string, ch <-chan int) {
	for {
		if i, ok := <-ch; ok {
			fmt.Printf("%s receiver, %v\n", name, i)
		} else {
			fmt.Printf("channel was closed. \n%s is done\n", name)
			break
		}
	}
}

/*
	・sender が初めに動き出す
	・受信側は buffer に queue が入るまで wait するので、そもそも送信先が定義されていない状況で、main thread で立ち上がると deadlock panic になる.
*/

func sender() {
	ch := make(chan int)
	go receiver("1st", ch)
	go receiver("2nd", ch)

	i := 0
	for i < 10 {
		fmt.Println("sender ", i)
		ch <- i
		i++
	}
	close(ch)
	time.Sleep(200 * time.Millisecond)
}
