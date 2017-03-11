// Copyright 2016 mikan. All rights reserved.

package main

import (
	"fmt"
	"unicode"
)

// reverse reverses an slice of bytes in place.
func reverse(s []byte) {

	// Bad answer. Use decodeRune and rotate

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] > unicode.MaxASCII {
			// exchange triple-byte rune
			// Limitation 1: other multi-byte runes are unsupported
			// Limitation 2: should be exchange the same byte length...
			s[i], s[j-2] = s[j-2], s[i]
			s[i+1], s[j-1] = s[j-1], s[i+1]
			s[i+2], s[j] = s[j], s[i+2]
			i += 2
			j -= 2
		} else {
			s[i], s[j] = s[j], s[i]
		}
	}
}

func main() {
	a := []byte("a-b-c-d-e")
	reverse(a[:])
	fmt.Printf("%v %v\n", a, string(a)) // e-d-c-b-a
	j := []byte("あ-い-う-え-お")
	reverse(j[:])
	fmt.Printf("%v %v\n", j, string(j)) // お-え-う-い-あ
}
