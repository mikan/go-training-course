// Copyright 2016 mikan. All rights reserved.

package main

import (
	"fmt"

	"strings"

	"github.com/mikan/libmikan/input"
)

const DECIMAL_POINT = "."
const DELIMITER = ","
const DELIMITER_N = 3

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	dot := strings.Index(s, DECIMAL_POINT)
	if n <= 3 || (dot >= 0 && dot <= DELIMITER_N) {
		return s
	}
	if s[0] == '-' && (n <= DELIMITER_N+1 || (dot >= 0 && dot <= DELIMITER_N+1)) {
		return s
	}
	if dot >= 0 && !strings.Contains(s, DELIMITER) {
		return comma(s[:dot-DELIMITER_N]) + DELIMITER + s[dot-DELIMITER_N:]
	}
	return comma(s[:n-DELIMITER_N]) + DELIMITER + s[n-DELIMITER_N:]
}

// 1111.000
// 1 111.000

// 1111
// 1 1

func main() {
	for {
		n := input.Word("Input number")
		if input.IsQuit(n) {
			return
		}
		fmt.Println(comma(n))
	}
}
