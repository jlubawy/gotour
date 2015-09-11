// https://tour.golang.org/methods/16

// Exercise: Images
// Remember the picture generator you wrote earlier? Let's write another one, but this time it will return an implementation of image.Image instead of a slice of data.
//
// Define your own Image type, implement the necessary methods, and call pic.ShowImage.
//
// Bounds should return a image.Rectangle, like image.Rect(0, 0, w, h).
//
// ColorModel should return color.RGBAModel.
//
// At should return a color; the value v in the last picture generator corresponds to color.RGBA{v, v, 255, 255} in this one.

package main

import (
	"fmt"
)

func Exercise10() {
	fmt.Println(`Code shown below must be run in browser to show picture:

	package main

	import (
		"image"
		"image/color"

		"code.google.com/p/go-tour/pic"
	)

	type Image struct {
		x0 int
		y0 int
		x1 int
		y1 int
	}

	func (img *Image) ColorModel() color.Model {
		return color.RGBAModel
	}

	func (img *Image) Bounds() image.Rectangle {
		return image.Rectangle{image.Point{img.x0, img.y0}, image.Point{img.x1, img.y1}}
	}

	func (img *Image) At(x, y int) color.Color {
		return color.RGBA{uint8(x ^ y), uint8(y ^ x), 255, 255}
	}

	func Exercise10() {
		m := &Image{0, 0, 500, 500}
		pic.ShowImage(m)
	}`)
}
