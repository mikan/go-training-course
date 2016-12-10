// Copyright 2016 mikan. All rights reserved.

package main

import "testing"

func TestRotate1(t *testing.T) {
	expected := [...]int{1}
	actual := expected
	rotate(actual[:])
	if expected != actual {
		t.Errorf("expected=%v actual=%v", expected, actual)
	}
}

func TestRotate3(t *testing.T) {
	expected := [...]int{1, 2, 3}
	actual := expected

	// take 1 of 3 (123 -> 312)
	rotate(actual[:])
	if expected == actual {
		t.Errorf("expected=%v actual=%v", expected, actual)
	}

	// take 2 of 3 (312 -> 231)
	rotate(actual[:])
	if expected == actual {
		t.Errorf("expected=%v actual=%v", expected, actual)
	}

	// take 3 of 3 (231 -> 123)
	rotate(actual[:])
	if expected != actual {
		t.Errorf("expected=%v actual=%v", expected, actual)
	}
}
