// Copyright 2016 mikan. All rights reserved.

package main

import (
	"fmt"
	"os"
)

// "foo" -> "fo"
func main() {
	if len(os.Args[1:]) == 0 {
		fmt.Printf("Usage: %v abc abc bcd bcd cde cde\n", os.Args[0])
		return
	}
	fmt.Println(shrink(os.Args[1:]))
}

func shrink(texts []string) []string {
	for i := 0; i < len(texts)-1; i++ {
		if texts[i] == texts[i+1] {
			texts = remove(texts, i+1)
			i-- // evaluate again
		}
	}
	return texts
}

func remove(slice []string, i int) []string {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}
