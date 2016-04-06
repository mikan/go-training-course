// Copyright 2016 mikan. All rights reserved.

package main

import (
	"testing"
	"strings"
)

func TestJoin(t *testing.T) {
	sep := ","
	checkEquals(t, strings.Join([]string{"a", "b", "c"}, sep), Join(sep, "a", "b", "c"))
	checkEquals(t, strings.Join([]string{"a"}, sep), Join(sep, "a"))
	checkEquals(t, strings.Join([]string{""}, ""), Join("", ""))
}

func checkEquals(t *testing.T, expected, actual string) {
	if expected != actual {
		t.Errorf("expected %s but got %s", expected, actual)
	}
}