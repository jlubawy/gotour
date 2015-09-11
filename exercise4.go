// https://tour.golang.org/moretypes/22

// Exercise: Fibonacci closure
// Let's have some fun with functions.
//
// Implement a fibonacci function that returns a function (a closure) that returns successive fibonacci numbers.

package main

import (
	"fmt"
)

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	var temp int

	current, next := 0, 1

	return func() int {
		temp = current
		current = next
		next = temp + next

		return temp
	}
}

func Exercise4() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
