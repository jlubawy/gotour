package main

import (
	"fmt"
	"os"
	"strconv"
)

var funcs = []func(){
	Exercise1,
	Exercise2,
	Exercise3,
	Exercise4,
	Exercise5,
	Exercise6,
	Exercise7,
	Exercise8,
	Exercise9,
	Exercise10,
	Exercise11,
	Exercise12,
}

func main() {
	if len(os.Args) != 2 {
		usage()
	}

	arg := os.Args[1]

	f, err := strconv.ParseUint(arg, 10, 64)
	if err != nil || f < 1 || f > uint64(len(funcs)) {
		fmt.Fprintf(os.Stderr, "Invalid exercise '%s'.\n\n", arg)
		usage()
	}

	funcs[f-1]()
}

func usage() {
	fmt.Fprintf(os.Stderr, `Usage:

	gotour <exercise %d-%d>
`, 1, len(funcs))
	os.Exit(1)
}
