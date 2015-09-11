// https://tour.golang.org/concurrency/7

// Exercise: Equivalent Binary Trees
// There can be many different binary trees with the same sequence of values stored at the leaves. For example, here are two binary trees storing the sequence 1, 1, 2, 3, 5, 8, 13.
//
//     [see picture online]
//
// A function to check whether two binary trees store the same sequence is quite complex in most languages. We'll use Go's concurrency and channels to write a simple solution.
//
// This example uses the tree package, which defines the type:
//
// type Tree struct {
//     Left  *Tree
//     Value int
//     Right *Tree
// }
//
// 1. Implement the Walk function.
//
// 2. Test the Walk function.
//
// The function tree.New(k) constructs a randomly-structured binary tree holding the values k, 2k, 3k, ..., 10k.
//
// Create a new channel ch and kick off the walker:
//
// go Walk(tree.New(1), ch)
// Then read and print 10 values from the channel. It should be the numbers 1, 2, 3, ..., 10.
//
// 3. Implement the Same function using Walk to determine whether t1 and t2 store the same values.
//
// 4. Test the Same function.
//
// Same(tree.New(1), tree.New(1)) should return true, and Same(tree.New(1), tree.New(2)) should return false.

package main

import (
	"fmt"

	"code.google.com/p/go-tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	// If this is a leaf node then send, close, and return
	if t.Left == nil && t.Right == nil {
		ch <- t.Value
		close(ch)
		return
	}

	// Else this is a node so try to walk left first
	if t.Left != nil {
		ch_left := make(chan int)
		go Walk(t.Left, ch_left)

		for left := range ch_left {
			ch <- left
		}
	}

	ch <- t.Value

	// Then try to walk right
	if t.Right != nil {
		ch_right := make(chan int)
		go Walk(t.Right, ch_right)

		for right := range ch_right {
			ch <- right
		}
	}

	// Close the channel
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	// Compare all values in tree until one doesn't match
	for v1 := range ch1 {
		v2 := <-ch2

		if v1 != v2 {
			return false
		}
	}

	return true
}

func Exercise11() {
	// Uncomment to test Walk

	// ch := make(chan int)
	// go Walk(tree.New(1), ch)
	//
	// for v := range ch {
	// 	fmt.Printf("%d\n", v)
	// }

	fmt.Printf("tree.New(1) == tree.New(1) is %t\n", Same(tree.New(1), tree.New(1)))
	fmt.Printf("tree.New(1) == tree.New(2) is %t\n", Same(tree.New(1), tree.New(2)))
}
