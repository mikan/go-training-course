// Copyright 2016 mikan. All rights reserved.

package main

import (
	"fmt"
	"strings"

	"github.com/mikan/util/input"
)

func anagram(s1, s2 string) bool {
	if s1 == s2 {
		return true
	}
	// Remove whitespace
	s1 = strings.Replace(s1, " ", "", -1)
	s2 = strings.Replace(s2, " ", "", -1)
	if len(s1) != len(s2) {
		return false
	}
	// Remove s1 characters with matched s2 characters
	for i := 0; i < len(s2); i++ {
		if !strings.Contains(s1, string(s2[i])) {
			return false // break
		}
		s1 = strings.Replace(s1, string(s2[i]), "", 1)
	}
	return len(s1) == 0 // Confirm removal result
}

func main() {
	for {
		s1 := input.SingleLine("Input s1")
		if input.IsQuit(s1) {
			return
		}
		s2 := input.SingleLine("Input s2")
		if input.IsQuit(s2) {
			return
		}
		fmt.Printf("anagram(s1, s2) = %t\n", anagram(s1, s2))
	}
}
