package main

import (
	"fmt"
	"golang.org/x/tour/tree"
	"sort"
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

func testWalk1() {
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
	close(ch) // len を指定無しにしてる場合など、close をして range を終了させる
}

func testWalk3() {
	ch := make(chan int, 10)
	ch2 := make(chan int, 10)
	Walk3(tree.New(1), ch)
	Walk3(tree.New(1), ch2)
	for i := range ch {
		fmt.Println(i)
	}
	for i := range ch2 {
		fmt.Println(i)
	}
}

// 全く同じ構造の判定には使える
// が構造は違くても同じ要素を含んでいる場合の判定には使えない...
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk3(t1, ch1)
	go Walk3(t2, ch2)
	for v1 := range ch1 {
		v2, ok := <-ch2
		if !ok || v1 != v2 {
			return false
		}
	}
	return true
}

func testSame() {
	println("same 1,1", Same(tree.New(1), tree.New(1)))
	println("same 1,2", Same(tree.New(1), tree.New(2)))
}

func SortChan(ch *chan int) []int {
	var (
		sortedSlice []int
	)
	for i := range *ch {
		sortedSlice = append(sortedSlice, i)
	}
	sort.Slice(sortedSlice, func(i, j int) bool { // i, j は index..
		return sortedSlice[i] < sortedSlice[j]
	})
	return sortedSlice
}

type intTuple struct {
	a, b int
}

func zip(a, b []int) ([]intTuple, error) {
	if len(a) != len(b) {
		return nil, fmt.Errorf("zip: arguments must be of same length")
	}

	r := make([]intTuple, len(a), len(a))
	for i, e := range a {
		r[i] = intTuple{e, b[i]}
	}
	return r, nil
}

func testSame2(slice1, slice2 []int) bool {
	zipped, err := zip(slice1, slice2)
	if err != nil {
		fmt.Println(err)
	}
	for _, pair := range zipped {
		if pair.a != pair.b {
			return false
		}
	}
	return true
}

func main() {
	//testSame()
	//SleepSort(&[]int{1, 5, 6, 3, 2, 8, 9, 10})

	ch := make(chan int)
	ch2 := make(chan int)
	go Walk3(tree.New(1), ch)
	go Walk3(tree.New(1), ch2)

	fmt.Println(testSame2(SortChan(&ch), SortChan(&ch2)))
}
