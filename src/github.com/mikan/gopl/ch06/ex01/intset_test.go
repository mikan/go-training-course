// Copyright 2016 mikan. All rights reserved.

package intset

import (
	"testing"
)

func TestLen(t *testing.T) {
	var set IntSet
	checkEquals(t, 0, set.Len())
	set.Add(1)
	checkEquals(t, 1, set.Len())
	set.Add(2)
	checkEquals(t, 2, set.Len())
	set.Add(3)
	checkEquals(t, 3, set.Len())
	if !(set.Has(1) && set.Has(2) && set.Has(3)) {
		t.Errorf("missing added element(s)")
	}
}

func TestRemove(t *testing.T) {
	var set IntSet
	set.Add(1)
	set.Add(2)
	set.Add(3)
	set.Remove(2)
	checkEquals(t, 2, set.Len()) // Contains 1 and 3
	if set.Has(2) {
		t.Errorf("element %v not removed", 2)
	}
	if !(set.Has(1) && set.Has(3)) {
		t.Errorf("missing added element(s)")
	}
}

func TestClear(t *testing.T) {
	var set IntSet
	checkEquals(t, 0, set.Len())
	set.Add(1)
	checkEquals(t, 1, set.Len())
	set.Clear()
	checkEquals(t, 0, set.Len())
}

func TestCopy(t *testing.T) {
	var set IntSet
	set.Add(1)
	set2 := set.Copy()
	if set.String() != set2.String() {
		t.Errorf("String() mismatch: %s <-> %s", set.String(), set2.String())
	}
	set.Add(2)
	set.Add(3)
	set2.Add(2)
	checkEquals(t, 3, set.Len())
	checkEquals(t, 2, set2.Len())
	if set.String() == set2.String() {
		t.Errorf("String() match unfortunately: %s <-> %s", set.String(), set2.String())
	}
}

func checkEquals(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("expected %v but got %v", expected, actual)
	}
}
