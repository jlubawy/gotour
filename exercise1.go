// https://tour.golang.org/flowcontrol/8

// Exercise: Loops and Functions
// As a simple way to play with functions and loops, implement the square root function using Newton's method.
//
// In this case, Newton's method is to approximate Sqrt(x) by picking a starting point z and then repeating:
//
// z = z - [(z^2 - x) / 2z]
//
// To begin with, just repeat that calculation 10 times and see how close you get to the answer for various values (1, 2, 3, ...).
//
// Next, change the loop condition to stop once the value has stopped changing (or only changes by a very small delta).
// See if that's more or fewer iterations. How close are you to the math.Sqrt?
//
// Hint: to declare and initialize a floating point value, give it floating point syntax or use a conversion:
//
// z := float64(1)
// z := 1.0

package main

import (
	"fmt"
	"math"
)

const (
	DELTA float64 = 0.1e-9
)

func sqrt(x float64) float64 {
	z := 1.0
	temp := 1.0

	for {
		temp = z

		z = z - ((z*z - x) / (2 * z))

		// Measure delta
		if math.Abs(z-temp) < DELTA {
			break
		}
	}

	return z
}

func Exercise1() {
	for i := 10.0; i <= 100.0; i += 10.0 {
		fmt.Printf("\nNewton's Method[%.0f] = %g", i, sqrt(i))
		fmt.Printf("\n   Math Package[%.0f] = %g", i, math.Sqrt(i))
		fmt.Printf("\n          Delta[%.0f] = %g", i, sqrt(i)-math.Sqrt(i))
	}
}
