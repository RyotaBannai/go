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

func tester() {
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

//func Same(t1, t2, *tree.Tree) bool {
//
//}

// 全部完了する前に close してしまう..
// sub tree が close してしまうため　-> children には chan を渡さないようにする → Walk3
func Walk2(t *tree.Tree, ch chan int) {
	if t != nil {
		ch <- t.Value
		Walk2(t.Left, ch)
		Walk2(t.Right, ch)
	}
	close(ch)
}

func Walk3(t *tree.Tree, ch chan int) {
	var subWalk func(*tree.Tree)
	subWalk = func(t *tree.Tree) {
		if t != nil {
			ch <- t.Value
			subWalk(t.Left)
			subWalk(t.Right)
		}
	}
	subWalk(t)
	close(ch)
}

func main() {
	ch := make(chan int, 10)
	go Walk3(tree.New(1), ch)
	for i := range ch {
		fmt.Println(i)
	}
}
