// Copyright 2016 mikan. All rights reserved.

package intset

import (
	"reflect"
	"testing"
)

func TestElms(t *testing.T) {
	var set IntSet
	set.Add(1)
	set.Add(2)
	set.Add(3)
	elms := set.Elms()
	if !reflect.DeepEqual(elms, []int{1, 2, 3}) {
		t.Errorf("slices mismatch: %v <-> %v", elms, []int{1, 2, 3})
	}
}
