package main

import (
	"fmt"
	"time"
)

// どの case も準備できていない場合でもブロックせずに送受信したい場合は、default を使う.

func selectDefault() {
	tick := time.Tick(1000 * time.Millisecond) // 定期実行
	boom := time.After(500 * time.Millisecond) // 初回のみ
	count := 1
	for count < 1000 {
		count++
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("boom.")
		default:
			fmt.Println("	.")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
