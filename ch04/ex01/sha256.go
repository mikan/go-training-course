// Copyright 2016 mikan. All rights reserved.

package main

import (
	"crypto/sha256"
	"fmt"

	"github.com/mikan/go-training-course/ch02/ex03/popcount"
)

func main() {
	v1 := textInput()
	v2 := textInput()
	c1 := sha256.Sum256([]byte(v1))
	c2 := sha256.Sum256([]byte(v2))
	fmt.Printf("popDiff(%v, %v): %v\n", v1, v2, popDiff(c1, c2))
}

func textInput() string {
	fmt.Print("input text > ")
	var text string
	fmt.Scan(&text)
	return text
}

func popDiff(c1, c2 [32]byte) int {
	count := 0
	for i := 0; i < 32; i++ {
		count += popcount.PopCount(uint64(c1[i]) ^ uint64(c2[i]))
	}
	return count
}
