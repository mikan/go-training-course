// Copyright 2016 mikan. All rights reserved.

package main

import (
	"fmt"
	"os"
	"unicode"
)

func main() {
	if len(os.Args[1:]) == 0 {
		fmt.Printf("Usage: %v foo\n", os.Args[0])
		return
	}
	v := []byte(os.Args[1])
	result := string(uToA(v))
	fmt.Printf("input:  %v\n", os.Args[1]) // input:  Hello,世界!
	fmt.Printf("result: %v\n", result)     // result: Hello,!
}

func uToA(text []byte) []byte {
	for i := 0; i < len(text); i++ {
		fmt.Printf("%x %b\n", text[i], text[i])
		if text[i] > unicode.MaxASCII {
			text = remove(text, i)
			i--
		}
	}
	return text
}

func remove(slice []byte, i int) []byte {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}
