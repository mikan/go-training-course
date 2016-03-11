// Copyright 2016 mikan. All rights reserved.

package main

import (
	"reflect"
	"testing"
)

func TestUToA(t *testing.T) {
	input := []byte("Hello,世界　　　　!")
	expected := []byte("Hello,世界 !")
	actual := uToA(input)
	if !reflect.DeepEqual(expected[:], actual) {
		t.Errorf("expected=%v actual=%v", expected, actual)
	}
}
