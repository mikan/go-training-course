// Copyright 2016 mikan. All rights reserved.

package main

import "testing"

func TestMax(t *testing.T) {
	checkEquals(t, -1, max(-1))
	checkEquals(t, 0, max(0))
	checkEquals(t, 1, max(1))
	checkEquals(t, 0, max(-10, -1, 0))
	checkEquals(t, 100, max(-10, -1, 0, 1, 10, 100))
	checkEquals(t, 1, max(1, 1, 1))
}

func TestMin(t *testing.T) {
	checkEquals(t, -1, min(-1))
	checkEquals(t, 0, min(0))
	checkEquals(t, 1, min(1))
	checkEquals(t, -10, min(-10, -1, 0))
	checkEquals(t, -10, min(-10, -1, 0, 1, 10, 100))
	checkEquals(t, 1, min(1, 1, 1))
}

func TestMax2(t *testing.T) {
	checkEquals(t, -1, max2(-1))
	checkEquals(t, 0, max2(0))
	checkEquals(t, 1, max2(1))
	checkEquals(t, 0, max2(-10, -1, 0))
	checkEquals(t, 100, max2(-10, -1, 0, 1, 10, 100))
	checkEquals(t, 1, max2(1, 1, 1))
}

func TestMin2(t *testing.T) {
	checkEquals(t, -1, min2(-1))
	checkEquals(t, 0, min2(0))
	checkEquals(t, 1, min2(1))
	checkEquals(t, -10, min2(-10, -1, 0))
	checkEquals(t, -10, min2(-10, -1, 0, 1, 10, 100))
	checkEquals(t, 1, min2(1, 1, 1))
}

func checkEquals(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("expected %s but got %s", expected, actual)
	}
}
