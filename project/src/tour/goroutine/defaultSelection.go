package main

import "time"

// どの case も準備できていない場合でもブロックせずに送受信したい場合は、default を使う.

func main() {
	tick:= time.Tick(100 * time.Millisecond)
	boom:= time.After(500 * time.Millisecond)
}