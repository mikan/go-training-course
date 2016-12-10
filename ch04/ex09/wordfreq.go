// Copyright 2016 mikan. All rights reserved.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Counts the words in a text file.
//
// Sample input in GOPATH:
// /LICENSE                CC-BY-NC-SA
// /README.md              The ReadMe
// /res/alice.txt          Alice's Adventures in Wonderland
// /res/war_and_peace.txt  War and Peace by Gutenberg
func main() {
	path := os.Getenv("GOPATH") + "/LICENSE" // Default is LICENSE
	if len(os.Args[1:]) > 0 {
		path = os.Args[1]
	}
	fmt.Printf("Loading %s ...\n", path)
	counts := make(map[string]int)
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	input := bufio.NewScanner(f)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		text := strings.ToLower(input.Text()) // get text with lower-case normalization
		if strings.HasSuffix(text, ".") {
			text = strings.Replace(text, ".", "", 1)
		}
		if strings.HasSuffix(text, ",") {
			text = strings.Replace(text, ",", "", 1)
		}
		counts[text]++
	}
	fmt.Println("count\tword")
	for word, count := range counts {
		fmt.Printf("%d\t%s\n", count, word)
	}
}

// Loading ./LICENSE ...
// count   word
// 1       commons
// 1       4.0
// 1       http://gopl.io
// 2       licensed
// 3       a
// 1       the
// 1       view
// 1       w
// 1       &
// 2       this
// 2       license
// 1       alan
// 1       gopl.io
// 1       kernighan
// 1       cc-by-nc-sa
// 1       to
// 1       copy
// 1       http://creativecommonsorg/licenses/by-nc-sa/4.0/.
// 1       visit
// 1       some
// 1       brian
// 1       is
// 1       international
// 1       of
// 1       based
// 1       on
// 1       authored
// 2       under
// 1       creative
// 1       are
// 1       by
// 1       donovan
// 1       work
// 1       attribution-noncommercial-sharealike
// 1       codes
