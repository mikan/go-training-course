// Copyright 2016 mikan. All rights reserved.

package main

import (
	"testing"
)

type inputTest struct {
	input string
	len   int
}

func TestByteCounter_Write(t *testing.T) {
	input := []inputTest{
		{"hello", 5},
		{"hello dolly", 11},
		{"hello\ndolly", 11},
	}
	for _, i := range input {
		var counter ByteCounter
		counter.Write([]byte(i.input))
		if int(counter) != i.len {
			t.Errorf("got %v want %v", counter, i.len)
		}
	}
}

func TestWordCounter_Write(t *testing.T) {
	input := []inputTest{
		{"hello", 1},
		{"hello dolly", 2},
		{"hello\ndolly", 2},
	}
	for _, i := range input {
		var counter WordCounter
		counter.Write([]byte(i.input))
		if int(counter) != i.len {
			t.Errorf("got %v want %v", counter, i.len)
		}
	}
}

func TestLineCounter_Write(t *testing.T) {
	input := []inputTest{
		{"hello", 1},
		{"hello dolly", 1},
		{"hello\ndolly", 2},
	}
	for _, i := range input {
		var counter LineCounter
		counter.Write([]byte(i.input))
		if int(counter) != i.len {
			t.Errorf("got %v want %v", counter, i.len)
		}
	}
}
