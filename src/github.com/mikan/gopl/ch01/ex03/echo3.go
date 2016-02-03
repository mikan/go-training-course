// Copyright 2015-2016 mikan. All rights reserved.

// Echo3 prints its command-line arguments.
package main
import (
	"fmt"
	"strings"
	"time"
)
func main() {
	input := []string{"foo", "bar", "baz"}
	secs1 := loop(1, 10000000, input)
	secs2 := loop(2, 10000000, input)
	fmt.Printf("run1: %.2fs\n", secs1);
	fmt.Printf("run2: %.2fs\n", secs2);
}

func run1(count []string) {
	s, sep := "", ""
	for _, arg := range count {
		s += sep + arg
		sep = " "
	}
	// fmt.Println(s)
}

func run2(count []string) {
	// fmt.Println(strings.Join(count, " "))
	strings.Join(count, " ")
}

func loop(method int, count int, input []string) float64 {
	start := time.Now()
	for i := 1; i < count; i++ {
		switch method {
		case 1: run1(input)
		case 2: run2(input)
		default: panic("unknown method")
		}
	}
	return time.Since(start).Seconds()
}

// Result (Intel Xeon X3320 / Windows 10):
// run1: 2.95s
// run2: 2.16s

