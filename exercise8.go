// https://tour.golang.org/methods/12

// Exercise: rot13Reader
// A common pattern is an io.Reader that wraps another io.Reader, modifying the stream in some way.
//
// For example, the gzip.NewReader function takes an io.Reader (a stream of compressed data) and returns a *gzip.Reader that also implements io.Reader (a stream of the decompressed data).
//
// Implement a rot13Reader that implements io.Reader and reads from an io.Reader, modifying the stream by applying the rot13 substitution cipher to all alphabetical characters.
//
// The rot13Reader type is provided for you. Make it an io.Reader by implementing its Read method.

package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot13 *rot13Reader) Read(b []byte) (int, error) {

	n, err := rot13.r.Read(b)

	for i := 0; i < n; i++ {
		if isUpper(b[i]) {
			b[i] = (((b[i] - 65) + 13) % 26) + 65
		} else if isLower(b[i]) {
			b[i] = (((b[i] - 97) + 13) % 26) + 97
		}
	}

	return n, err
}

func isUpper(c byte) bool {
	return ((c >= 65) && (c <= 90))
}

func isLower(c byte) bool {
	return ((c >= 97) && (c <= 122))
}

func Exercise8() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
