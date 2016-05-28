// Copyright 2016 mikan. All rights reserved.

package main

import (
	"io/ioutil"
	"testing"
)

func TestCountingWriter(t *testing.T) {
	w, c := CountingWriter(ioutil.Discard)
	input := "Hoge"
	n1, _ := w.Write([]byte(input))
	n2, _ := w.Write([]byte(input))
	n3, _ := w.Write([]byte(input))
	if *c != int64(n1+n2+n3) {
		t.Errorf("got %v want %v", *c, n1+n2+n3)
	}
}
