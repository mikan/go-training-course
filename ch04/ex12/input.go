package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Word gets input a word from stdin.
func Word(msg string) string {
	fmt.Print(msg + " > ")
	var text string
	fmt.Scan(&text)
	return Normalize(text)
}

// Normalize returns lower cased & spaces removed string of the input.
func Normalize(msg string) string {
	msg = strings.Trim(msg, " ")
	msg = strings.ToLower(msg)
	return msg
}

// SingleLine gets input a line from stdin.
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

// IsQuit returns whether command is quit (or exit).
func IsQuit(msg string) bool {
	text := Normalize(msg)
	return text == "quit" || text == "exit" || text == "42"
}
