// Copyright 2016 mikan. All rights reserved.

// Utilities for common text input operations.
package input

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// Input a word from stdin.
func Word(msg string) string {
	fmt.Print(msg + " > ")
	var text string
	fmt.Scan(&text)
	return Normalize(text)
}

// Input a line from stdin.
func SingleLine(msg string) string {
	fmt.Print(msg + " > ")
	in := bufio.NewReader(os.Stdin)
	line, _, err := in.ReadLine() // returns line, isPrefix, error
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	return string(line)
}

// Input multiple lines from stdin.
func MultiLine(msg string) string {
	fmt.Println(msg + " (Ctrl+D to complete) >>>")
	var body string
	in := bufio.NewReader(os.Stdin)
	for {
		line, _, err := in.ReadLine() // returns line, isPrefix, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		body += string(line) + "\n"
	}
	return body
}

// Returns lower cased & spaces removed string of the input.
func Normalize(msg string) string {
	msg = strings.Trim(msg, " ")
	msg = strings.ToLower(msg)
	return msg
}

// Returns whether command is quit (or exit).
func IsQuit(msg string) bool {
	text := Normalize(msg)
	return text == "quit" || text == "exit" || text == "42"
}
