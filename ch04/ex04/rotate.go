// Copyright 2016 mikan. All rights reserved.

package main

import "fmt"

func rotate(s []int) {
	tail := len(s) - 1
	for i := 0; i < len(s); i++ { // O(n)
		s[i], s[tail] = s[tail], s[i] // exchange i and tail
	}
}

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	fmt.Println(a) // [0 1 2 3 4 5]
	rotate(a[:])
	fmt.Println(a) // [5 0 1 2 3 4]
	rotate(a[:])
	fmt.Println(a) // [4 5 0 1 2 3]
	rotate(a[:])
	fmt.Println(a) // [3 4 5 0 1 2]
	rotate(a[:])
	fmt.Println(a) // [2 3 4 5 0 1]
	rotate(a[:])
	fmt.Println(a) // [1 2 3 4 5 0]
	rotate(a[:])
	fmt.Println(a) // [0 1 2 3 4 5]
}
