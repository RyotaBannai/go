package main

import (
	"fmt"
	"golang.org/x/tour/tree"
	"time"
)

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

// Walk walks the tree t sending all values from the tree to the channel.
func Walk(t *tree.Tree, ch chan int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("This branch is empty.")
		}
	}()
	ch <- t.Value
	go Walk(t.Left, ch)
	go Walk(t.Right, ch)
}

//func Same(t1, t2, *tree.Tree) bool {
//
//}

func main() {
	ch := make(chan int)
	t1 := tree.New(1)
	go Walk(t1, ch)
	isEnd := time.After(5 * time.Second) // type is chan
	var values []int
	for {
		select {
		case <-isEnd:
			fmt.Println(values)
			return
		case value := <-ch:
			values = append(values, value)
		}
	}
}
