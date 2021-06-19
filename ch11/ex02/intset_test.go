// Copyright 2021 mikan. All rights reserved.

package intset

import "testing"

func TestIntSet(t *testing.T) {
	tests := []struct {
		adds []int
	}{
		{[]int{}},
		{[]int{1}},
		{[]int{1, 2}},
		{[]int{1, 2, 3}},
	}
	for i, test := range tests {
		var tgt IntSet
		std := make(map[int]struct{})
		for _, add := range test.adds {
			tgt.Add(add)
			std[add] = struct{}{}
		}
		for add := range std {
			has :=  tgt.Has(add)
			if !has {
				t.Errorf("#%d Has(%d) = false, expected true", i, add)
			}
		}
	}
}
