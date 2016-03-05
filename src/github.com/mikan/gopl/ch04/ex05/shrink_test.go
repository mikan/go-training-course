// Copyright 2015-2016 mikan. All rights reserved.

package main

import (
	"reflect"
	"testing"
)

func TestShrink(t *testing.T) {
	input := [...]string{"ABC", "ABC", "ABC", "BCD", "ABC"}
	expected := [...]string{"ABC", "BCD", "ABC"}
	actual := shrink(input[:])
	if !reflect.DeepEqual(expected[:], actual) {
		t.Errorf("expected=%v actual=%v", expected, actual)
	}
}
