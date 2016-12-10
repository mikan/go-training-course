// Copyright 2015-2016 mikan. All rights reserved.

// Dup2 prints the count and text of lines that appear more than once in the input.
// It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)           // key: line, value: count
	sourceFiles := make(map[string][]string) // key: line, value: list of file names
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, sourceFiles)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, sourceFiles)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t[%s]\n", n, line, strings.Join(sourceFiles[line], ","))
		}
	}
}

func countLines(f *os.File, counts map[string]int, sourceFiles map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		sourceFiles[input.Text()] = append(sourceFiles[input.Text()], f.Name())
	}
	// NOTE: ignoring potential errors from input.Err()
}
