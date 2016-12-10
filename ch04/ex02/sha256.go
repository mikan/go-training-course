// Copyright 2016 mikan. All rights reserved.

package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var length = 256
	if len(os.Args[1:]) == 1 {
		l, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Printf("Usage: %v {256,384,512}\n", os.Args[0])
			return
		}
		length = l
	}
	fmt.Printf("SHA%v: %x\n", length, makeSHA(textInput(), length))
}

func textInput() string {
	fmt.Print("input text > ")
	var text string
	fmt.Scan(&text)
	return text
}

func makeSHA(text string, length int) []byte {
	switch length {
	case 512:
		v := sha512.Sum512([]byte(text))
		return v[:]
	case 384:
		v := sha512.Sum384([]byte(text))
		return v[:]
	case 256:
		fallthrough
	default:
		v := sha256.Sum256([]byte(text))
		return v[:]
	}
}
