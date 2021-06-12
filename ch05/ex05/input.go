package main

import (
	"fmt"
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
