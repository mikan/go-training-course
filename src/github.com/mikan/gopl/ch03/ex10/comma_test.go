// Copyright 2016 mikan. All rights reserved.
package main

import "testing"

func TestComma(t *testing.T) {
	verify(t, "", comma(""))
	verify(t, "1", comma("1"))
	verify(t, "11", comma("11"))
	verify(t, "111", comma("111"))
	verify(t, "1,111", comma("1111"))
	verify(t, "11,111", comma("11111"))
	verify(t, "111,111", comma("111111"))
	verify(t, "1,111,111", comma("1111111"))
	verify(t, "-1", comma("-1"))
	verify(t, "-11", comma("-11"))
	verify(t, "-111", comma("-111"))
	verify(t, "-1,111", comma("-1111"))
	verify(t, "-11,111", comma("-11111"))
	verify(t, "-111,111", comma("-111111"))
	verify(t, "-1,111,111", comma("-1111111"))
}

func verify(t *testing.T, expected, actual string) {
	if expected != actual {
		t.Errorf("expected %s but got %s", expected, actual)
	}
}
