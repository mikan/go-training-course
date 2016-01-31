// Copyright 2015-2016 mikan. All rights reserved.

// Echo3 prints its command-line arguments.
package main
import (
	"fmt"
	"strings"
)
func main() {
	input := []string{"foo", "bar", "baz"}
	run1(input)
	run2(input)
	// TODO: パフォーマンス計測コードを書く
}

func run1(count []string) {
	s, sep := "", ""
	for _, arg := range count {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func run2(count []string) {
	fmt.Println(strings.Join(count, " "))
}