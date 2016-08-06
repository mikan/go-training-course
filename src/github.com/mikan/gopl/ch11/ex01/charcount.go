// Copyright 2015-2016 mikan. All rights reserved.

// Charcount computes counts of Unicode characters.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	count(os.Stdin)
}

func count(reader io.Reader) {
	counts := make(map[rune]int)    // counts of Unicode characters
	var utfLen [utf8.UTFMax + 1]int // count of length of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters

	fmt.Println("Input characters (Ctrl-D to end) >>>")
	in := bufio.NewReader(reader)
	for {
		r, n, err := in.ReadRune() // returns rune, nByets, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utfLen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utfLen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
