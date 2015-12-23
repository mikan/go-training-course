// Copyright 2015-2016 mikan. All rights reserved.

// Echo2 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + strconv.Itoa(i) + " " + os.Args[i]
		sep = "\n"
	}
	fmt.Println(s)
}

// Result:
// $ ./bin/ex02 a b c
// 1 a
// 2 b
// 3 c
