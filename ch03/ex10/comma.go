// Copyright 2016 mikan. All rights reserved.

package main

import (
	"bytes"
	"fmt"

	"strings"

	"github.com/mikan/libmikan/input"
)

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	var buf bytes.Buffer
	i := 0
	buf.WriteByte(s[i])
	i++
	if strings.HasPrefix(s, "-") {
		buf.WriteByte(s[i])
		i++
	}
	for ; i < len(s); i++ {
		if (len(s)-i)%3 == 0 {
			buf.WriteByte(',')
		}
		buf.WriteByte(s[i])
	}
	return buf.String()
}

func main() {
	for {
		n := input.Word("Input number")
		if input.IsQuit(n) {
			return
		}
		fmt.Println(comma(n))
	}
}
