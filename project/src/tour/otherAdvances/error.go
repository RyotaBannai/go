package main

import (
	"time"
	"fmt"
)

type MyError struct {
	When time.Time
	What string
}

// Errors: myError 型
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

// myError 型
func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	// err == nil なら成功であると考える
	// THINK: run() から myError 以外の戻り値が欲しい時、はどう myError を実装すれば良いんだろうか... -> 返却値を error と value のタプルにする
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
