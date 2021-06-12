// Copyright 2016 mikan. All rights reserved.

package main

import "regexp"

func main() {
	for {
		text := SingleLine("Text")
		if IsQuit(text) {
			break
		}
		println(expand(text, func(s string) string { return "/* " + s + " */" }))
	}
}

func expand(s string, f func(string) string) string {
	r := regexp.MustCompile(`\$\w+`)                                            // \w = a word character
	return r.ReplaceAllStringFunc(s, func(s string) string { return f(s[1:]) }) // remove "$"
}
