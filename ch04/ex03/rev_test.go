// Copyright 2016 mikan. All rights reserved.

package main

import "testing"

func TestReverse6(t *testing.T) {
	expected := [...]int{0, 1, 2, 3, 4, 5}
	actual := expected
	reverse6(&actual)
	if expected == actual {
		t.Errorf("unchanged in reverse6() %v -> %v", expected, actual)
	}
	reverse6(&actual)
	if expected != actual {
		t.Errorf("expected=%v actual=%v", expected, actual)
	}
}

func TestReverse32(t *testing.T) {
	input := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var expected [32]int
	copy(expected[:], input[:])
	actual := expected
	reverse32(&actual)
	if expected == actual {
		t.Errorf("unchanged in reverse6() %v -> %v", expected, actual)
	}
	reverse32(&actual)
	if expected != actual {
		t.Errorf("expected=%v actual=%v", expected, actual)
	}
}
