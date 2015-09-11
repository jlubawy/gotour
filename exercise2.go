// https://tour.golang.org/moretypes/14

// Exercise: Slices
// Implement Pic. It should return a slice of length dy, each element of which is a slice of dx 8-bit unsigned integers.
// When you run the program, it will display your picture, interpreting the integers as grayscale (well, bluescale) values.
//
// The choice of image is up to you. Interesting functions include (x+y)/2, x*y, and x^y.
//
// (You need to use a loop to allocate each []uint8 inside the [][]uint8.)
//
// (Use uint8(intValue) to convert between types.)

package main

import (
	"fmt"
)

func Exercise2() {
	fmt.Println(`Code shown below must be run in browser to show picture:

	package main

	import (
		"code.google.com/p/go-tour/pic"
	)

	func Pic(dx, dy int) [][]uint8 {

		// Create slice of length dy
		picture := make([][]uint8, dy)

		// Each element which is a slice of length dx
		for i := range picture {
			picture[i] = make([]uint8, dx)
		}

		for y := 0; y < len(picture); y++ {
			for x := 0; x < len(picture[y]); x++ {
				picture[y][x] = uint8(x ^ y)
			}
		}

		return picture
	}

	func main() {
		pic.Show(Pic)
	}`)
}
