// https://tour.golang.org/methods/11

// Exercise: Readers
// Implement a Reader type that emits an infinite stream of the ASCII character 'A'.

package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (r MyReader) Read(b []byte) (int, error) {
	for i := 0; i < len(b); i++ {
		b[i] = byte('A')
	}

	return len(b), nil
}

func Exercise7() {
	reader.Validate(MyReader{})
}
