// Copyright 2016 mikan. All rights reserved.

package intset

import (
	"testing"
)

func TestAddAll(t *testing.T) {
	var set IntSet
	set.AddAll(1, 2, 3)
	if !(set.Has(1) && set.Has(2) && set.Has(3)) {
		t.Error("missing added element(s)")
	}
	set.AddAll(4, 5, 6)
	if !(set.Has(4) && set.Has(5) && set.Has(6)) {
		t.Error("missing added element(s)")
	}
}
