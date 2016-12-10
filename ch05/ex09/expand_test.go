// Copyright 2016 mikan. All rights reserved.

package main

import "testing"

func TestExpand(t *testing.T) {
	verify(t, "", expand("", f))
	verify(t, "foo", expand("foo", f))
	verify(t, "[foo]", expand("$foo", f))
	verify(t, "[foo][bar][baz]", expand("$foo$bar$baz", f))
	verify(t, "[foo] [bar] [baz]", expand("$foo $bar $baz", f))
	verify(t, "1 [foo] 2 [bar] 3 [baz] 4", expand("1 $foo 2 $bar 3 $baz 4", f))
}

func verify(t *testing.T, expected, actual string) {
	if expected != actual {
		t.Errorf("expected %s but got %s", expected, actual)
	}
}

func f(s string) string {
	return "[" + s + "]"
}
