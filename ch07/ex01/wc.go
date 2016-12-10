// Copyright 2016 mikan. All rights reserved.

package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // convert int to ByteCounter
	return len(p), nil
}

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(string(p)))
	scanner.Split(bufio.ScanWords)
	var counts int
	for scanner.Scan() {
		scanner.Text()
		counts++
	}
	*c += WordCounter(counts) // convert int to WordCounter
	return counts, nil
}

type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(string(p)))
	scanner.Split(bufio.ScanLines)
	var counts int
	for scanner.Scan() {
		scanner.Text()
		counts++
	}
	*c += LineCounter(counts) // convert int to LineCounter
	return counts, nil
}

var c *bool = flag.Bool("c", false, "print the byte counts")
var w *bool = flag.Bool("w", false, "print the word counts")
var l *bool = flag.Bool("l", false, "print the newline counts")

// main returns counts of bytes, words and lines of specified file.
func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		flag.Usage()
		return
	}
	var tBytes ByteCounter
	var tWords WordCounter
	var tLines LineCounter
	for i := range flag.Args() {
		file := flag.Arg(i)
		f, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		var bytes ByteCounter
		var words WordCounter
		var lines LineCounter
		bytes.Write(f)
		words.Write(f)
		lines.Write(f)
		tBytes += bytes
		tWords += words
		tLines += lines
		if *c == *w && *w == *l {
			fmt.Printf("%6d %6d %6d %s\n", lines, words, bytes, file)
		} else if *c {
			fmt.Printf("%6d %s\n", bytes, file)
		} else if *w {
			fmt.Printf("%6d %s\n", words, file)
		} else if *l {
			fmt.Printf("%6d %s\n", lines, file)
		}
	}
	if flag.NArg() > 1 {
		if *c == *w && *w == *l {
			fmt.Printf("%6d %6d %6d total\n", tLines, tWords, tBytes)
		} else if *c {
			fmt.Printf("%6d total\n", tBytes)
		} else if *w {
			fmt.Printf("%6d total\n", tWords)
		} else if *l {
			fmt.Printf("%6d total\n", tLines)
		}
	}
}
