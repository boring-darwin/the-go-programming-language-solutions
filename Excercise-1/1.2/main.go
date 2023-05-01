package main

import (
	"fmt"
	"os"
)

// Exercise 1.2: Modify the echo program to print the index and value of each of its arguments,
// one per line.

func main() {
	for i, s := range os.Args {
		fmt.Printf("%d %s\n", i, s)
	}
}
