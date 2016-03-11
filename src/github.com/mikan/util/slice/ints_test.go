// Copyright 2016 mikan. All rights reserved.

package slice
import "testing"

func TestAdd(t *testing.T) {
	var s IntSlice
	s.Add(1)
	s.Add(2)
	if s.ToSlice()[0] != 1 {
		t.Errorf("expected %s but got %s", 1, s.ToSlice()[1])
	}
	if s.ToSlice()[1] != 2 {
		t.Errorf("expected %s but got %s", 2, s.ToSlice()[2])
	}
}