// Copyright 2016 mikan. All rights reserved.

package main

import (
	"strings"
	"testing"
)

type inputTest struct {
	input string
	len   int
}

func TestLimitedReader_Read(t *testing.T) {
	limit := 10
	input := []inputTest{
		{"hello", 5},
		{"hello doll", limit},    // 10
		{"hello dolly", limit},   // 11 -> 10
		{"hello\ndolly!", limit}, // 12 -> 10
	}
	for _, i := range input {
		reader := LimitReader(strings.NewReader(i.input), int64(limit))
		n, err := reader.Read([]byte(i.input))
		if err != nil {
			t.Errorf("got %v", err)
		}
		if n != i.len {
			t.Errorf("got %v want %v", n, i.len)
		}
	}
}
