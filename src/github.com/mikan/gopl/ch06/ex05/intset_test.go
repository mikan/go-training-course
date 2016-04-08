// Copyright 2016 mikan. All rights reserved.

package intset

import (
	"reflect"
	"testing"
)

func TestElms(t *testing.T) {
	var set IntSet
	set.Add(1)
	set.Add(1024)
	set.Add(2048)
	elms := set.Elms()
	if !reflect.DeepEqual(elms, []int{1, 1024, 2048}) {
		t.Errorf("slices mismatch: %v <-> %v", elms, []int{1, 1024, 2048})
	}
}
