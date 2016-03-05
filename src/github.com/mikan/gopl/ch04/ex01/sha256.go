// Copyright 2016 mikan. All rights reserved.

package main

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"math"

	"github.com/mikan/gopl/ch02/ex03/popcount"
)

func main() {
	v1 := textInput()
	v2 := textInput()
	c1 := sha256.Sum256([]byte(v1))
	c2 := sha256.Sum256([]byte(v2))
	fmt.Printf("popDiff(%v, %v): %v\n", v1, v2, popDiff(c1[:], c2[:]))
}

func textInput() string {
	fmt.Printf("input text > ")
	var text string
	fmt.Scan(&text)
	return text
}

// 多分題意と違う希ガス
func popDiff(c1, c2 []byte) int {
	u1, _ := binary.Uvarint(c1)
	u2, _ := binary.Uvarint(c2)
	v1 := popcount.PopCount(u1)
	v2 := popcount.PopCount(u2)
	return int(math.Abs(float64(v1) - float64(v2)))
}
