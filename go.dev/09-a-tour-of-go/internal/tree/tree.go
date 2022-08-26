package tree

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int, top bool) {
	if t.Left != nil {
		Walk(t.Left, ch, false)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch, false)
	}
	if top {
		close(ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1, true)
	go Walk(t2, ch2, true)

	for x := range ch1 {
		if x != <-ch2 {
			return false
		}
	}
	return true
}

func Run() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
