// Copyright 2016 mikan. All rights reserved.

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
	counts := make(map[rune]int)    // counts of Unicode characters
	types := make(map[string]int)   // counts of character types
	var utfLen [utf8.UTFMax + 1]int // count of length of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters

	fmt.Println("MEMO: Use [Ctrl]+[D] to complete input.")
	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nBytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "chrcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utfLen[n]++

		// Count character types
		if unicode.IsLetter(r) {
			types["letter"]++
		}
		if unicode.IsControl(r) {
			types["control"]++
		}
		if unicode.IsDigit(r) {
			types["digit"]++
		}
		if unicode.IsGraphic(r) {
			types["graphic"]++
		}
		if unicode.IsLower(r) {
			types["lower"]++
		}
		if unicode.IsMark(r) {
			types["mark"]++
		}
		if unicode.IsNumber(r) {
			types["number"]++
		}
		if unicode.IsPrint(r) {
			types["print"]++
		}
		if unicode.IsPunct(r) {
			types["punct"]++
		}
		if unicode.IsSpace(r) {
			types["space"]++
		}
		if unicode.IsSymbol(r) {
			types["symbol"]++
		}
		if unicode.IsTitle(r) {
			types["title"]++
		}
		if unicode.IsUpper(r) {
			types["upper"]++
		}
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
	fmt.Print("\ntype\tcount\n")
	for c, n := range types {
		fmt.Printf("%s\t%d\n", c, n)
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}

// Sample input: [aA あ⊿]
//
// rune      count
// 'a'     1
// 'A'     1
// ' '     1
// 'あ'     1
// '⊿'     1
//
// len     count
// 1       3
// 2       0
// 3       2
// 4       0
//
// type    count
// letter  3
// graphic 5
// lower   1
// print   5
// upper   1
// space   1
// symbol  1
