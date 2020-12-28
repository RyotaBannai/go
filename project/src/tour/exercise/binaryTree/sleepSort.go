package main

import (
	"fmt"
	"time"
)

/*
	sleep sort:
		時間差で単純に append していくだけで実際に値のソートを行わない手法

		https://qiita.com/exotic-toybox/items/2ecd62d3fd32bdd5511b

	・goroutine 内で変数を使うときの注意点, go vet でチェック.
		https://qiita.com/sudix/items/67d4cad08fe88dcb9a6d
*/

func SleepSort(numList *[]int) {
	var (
		sortedList = make([]int, 0)
		start      = make(chan int)
	)
	fmt.Println(numList)
	/*
		channel を使用して並行処理を一斉開始
	*/
	for _, num := range *numList {
		go func(n int) {
			<-start
			time.Sleep(time.Duration(n) * time.Millisecond)
			sortedList = append(sortedList, n)
		}(num)
	}
	close(start) // channel を close するのはお作法.

	for {
		if len(*numList) == len(sortedList) {
			fmt.Printf("sortedList:\t%d\n", sortedList)
			return
		}
	}
}
