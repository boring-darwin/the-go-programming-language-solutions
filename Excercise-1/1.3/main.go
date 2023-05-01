package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// Exercise 1.3: Experiment to measure the difference in running time between our potentially
// inefficient versions and the one that uses strings.Join. (Section 1.6 illustrates part of the
// time package, and Section 11.4 shows how to write benchmark tests for systematic per-
// formance evaluation.)

func main() {
	t := time.Now()
	contact()
	totalT1 := time.Since(t)

	t = time.Now()
	join()
	totalT2 := time.Since(t)

	fmt.Printf("Time took for concat %d\nTime took for join %d\n", totalT1, totalT2)
}

func contact() string {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	return s
}

func join() string {
	return strings.Join(os.Args[1:], " ")
}
