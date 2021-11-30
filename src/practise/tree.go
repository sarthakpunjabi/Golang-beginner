package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

func Same(t1, t2 *tree.Tree) bool {
	treeCh1 := make(chan int)
	treeCh2 := make(chan int)
	go func() {
		Walk(t1, treeCh1)
		close(treeCh1)
	}()
	go func() {
		Walk(t2, treeCh2)
		close(treeCh2)
	}()

	for {
		val1 := <-treeCh1
		val2 := <-treeCh2
		if val1 == val2 {
			return true
		} else {
			return false
		}
	}
}

func main() {
	fmt.Println("Testing Walk function")
	ch := make(chan int)
	go func() {
		Walk(tree.New(1), ch)
		close(ch)
	}()
	for val := range ch {
		fmt.Println("Received from ch, ", val)
	}

	fmt.Println("Testing Same Function")
	fmt.Println("Test Case1")
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println("Test Case2")
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
