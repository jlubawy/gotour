// https://tour.golang.org/moretypes/19

// Exercise: Maps
// Implement WordCount. It should return a map of the counts of each “word” in the string s. The wc.Test function runs a test suite against the provided function and prints success or failure.
//
// You might find strings.Fields helpful.

package main

import (
	"strings"

	"code.google.com/p/go-tour/wc"
)

func wordCount(s string) map[string]int {
	words := strings.Fields(s)
	wc := make(map[string]int, len(words))

	for _, word := range words {
		wc[word] += 1
	}

	return wc
}

func Exercise3() {
	wc.Test(wordCount)
}
