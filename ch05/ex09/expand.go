// Copyright 2016 mikan. All rights reserved.

package main

import (
	"regexp"

	"github.com/mikan/libmikan/input"
)

func main() {
	for {
		text := input.SingleLine("Text")
		if input.IsQuit(text) {
			break
		}
		println(expand(text, func(s string) string { return "/* " + s + " */" }))
	}
}

func expand(s string, f func(string) string) string {
	r := regexp.MustCompile(`\$\w+`)                                            // \w = a word character
	return r.ReplaceAllStringFunc(s, func(s string) string { return f(s[1:]) }) // remove "$"
}
