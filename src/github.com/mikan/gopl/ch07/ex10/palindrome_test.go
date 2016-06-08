// Copyright 2016 mikan. All rights reserved.

package main

import "testing"

type inputTest struct {
	text     string
	expected bool
}

func TestIsPalindrome(t *testing.T) {
	input := []inputTest{
		{"ABBA", true},
		{"ABBB", false},
	}
	for _, i := range input {
		got := IsPalindrome(palindrome([]rune(i.text)))
		if got != i.expected {
			t.Errorf("got %v want %v (text=%s)", got, i.expected, i.text)
		}
	}
}
