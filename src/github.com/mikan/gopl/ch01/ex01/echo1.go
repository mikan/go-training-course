// Copyright 2015-2016 mikan. All rights reserved.

// Echo1 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Program name is " + os.Args[0])    // Added
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

// Result:
// $ ./bin/ex01 a b c
// Program name is ./bin/ex01
// a b c
