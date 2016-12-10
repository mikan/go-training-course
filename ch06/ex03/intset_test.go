// Copyright 2016 mikan. All rights reserved.

package intset

import (
	"testing"
)

func TestUnionWith(t *testing.T) {
	var a, b IntSet
	a.Add(1)
	a.Add(2)
	b.Add(2)
	b.Add(3)
	a.UnionWith(&b) // (1 2) | (2 3) = (1 2 3)
	if !(a.Has(1) && a.Has(2) && a.Has(3)) {
		t.Error("unexpected content: ", a.String())
	}
}

func TestIntersectWith(t *testing.T) {
	var a, b IntSet
	a.Add(1)
	a.Add(2)
	b.Add(2)
	b.Add(3)
	a.IntersectWith(&b) // (1 2) & (2 3) = (2)
	if !(a.Has(2)) {
		t.Error("unexpected content: ", a.String())
	}
	if a.Has(1) || a.Has(3) {
		t.Error("unexpected content: ", a.String())
	}
}

func TestSymmetricDifference(t *testing.T) {
	var a, b IntSet
	a.Add(1)
	a.Add(2)
	b.Add(2)
	b.Add(3)
	a.SymmetricDifference(&b) // (1 2) ^ (2 3) = (1 3)
	if !(a.Has(1) && a.Has(3)) {
		t.Error("unexpected content: ", a.String())
	}
	if a.Has(2) {
		t.Error("unexpected content: ", a.String())
	}
}
