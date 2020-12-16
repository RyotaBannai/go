package main

import (
	"time"
	"fmt"
)

type MyError struct {
	When time.Time
	What string
}

// Errors: error 型
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

// error 型
func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	// err == nil なら成功であると考える
	// THINK: run() から error 以外の戻り値が欲しい時、はどう error を実装すれば良いんだろうか...
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
