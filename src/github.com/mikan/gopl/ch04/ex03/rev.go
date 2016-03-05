// Copyright 2016 mikan. All rights reserved.

package main

import "fmt"

// reverse6 reverses an array of ints in place.
func reverse6(s *[6]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func reverse32(s *[32]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse6(&a)
	fmt.Println(a) // [5 4 3 2 1 0]
	var a32 [32]int
	copy(a32[:], a[:])
	reverse32(&a32)
	fmt.Println(a32) // [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 1 2 3 4 5]
}
