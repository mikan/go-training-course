// Copyright 2016 mikan. All rights reserved.

package main

import (
	"fmt"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	if len(os.Args[1:]) == 0 {
		fmt.Printf("Usage: %v foo\n", os.Args[0])
		return
	}
	v := []byte(os.Args[1])
	result := string(uToA(v))
	fmt.Printf("input:  %v\n", os.Args[1]) // input:  Hello, 世界　    　!
	fmt.Printf("result: %v\n", result)     // result: Hello, 世界 !
}

func uToA(text []byte) []byte {
	for i := 0; i < len(text) - 1; i++ {
		c := text[i:i+charLength(text[i])]
		if isSpace(c) {
			for j := 0; j < len(c);j++ {
				text[i+j] = ' ' // ascii space
			}
		}
	}
	for i := 0; i < len(text) - 1; i++ {
		if text[i] == ' ' && text[i+1] == ' ' {
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

func charLength(b byte) int {
	if b < 128 { // 10000000
		return 1
	}
	if b < 224 { // 11100000
		return 2
	}
	if b < 240 { // 11110000
		return 3
	}
	return 4
}

func isSpace(b []byte) bool {
	r,_ := utf8.DecodeRune(b)
	return unicode.IsSpace(r)
}