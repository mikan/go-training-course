// Copyright 2016 mikan. All rights reserved.
package main
import "testing"

func TestAnagram(t *testing.T) {
	verify(t, true, anagram("", ""))
	verify(t, true, anagram("a", "a"))
	verify(t, true, anagram("ab", "ba"))
	verify(t, false, anagram("a", "aa"))
	verify(t, false, anagram("ab", "aa"))
	verify(t, true, anagram("abc", "a cb"))
}

func verify(t *testing.T, expected, actual bool) {
	if expected != actual {
		t.Errorf("expected %s but got %s", expected, actual)
	}
}