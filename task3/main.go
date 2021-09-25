package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

//  Ну а почему бы и нет
func getSize() int {
	return 10
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	walkTree(t, ch)
	close(ch)
}

func walkTree(t *tree.Tree, ch chan int) {
	if t != nil {
		walkTree(t.Left, ch)
		ch <- t.Value
		walkTree(t.Right, ch)

	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int, getSize())
	ch2 := make(chan int, getSize())

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for {
		value1, ok1 := <-ch1
		value2, ok2 := <-ch2

		if value1 != value2 || ok1 != ok2 {
			return false
		}

		if !ok1 {
			break
		}
	}
	return true
}

func main() {

	ch := make(chan int, getSize())

	go Walk(tree.New(getSize()), ch)

	fmt.Println(Same(tree.New(1), tree.New(2)))
	fmt.Println(Same(tree.New(1), tree.New(1)))
}
