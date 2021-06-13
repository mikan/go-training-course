// Copyright 2015-2016 mikan. All rights reserved.

package main

import "testing"

var sampleInput = []string{"foo1", "foo2", "foo3", "foo4", "foo5", "foo6"}

func BenchmarkRun1(b *testing.B) {
	for i := 1; i < b.N; i++ {
		run1(sampleInput)
	}
}

func BenchmarkRun2(b *testing.B) {
	for i := 1; i < b.N; i++ {
		run2(sampleInput)
	}
}

// Result:
// go test github.com/mikan/gopl/ex01/ex03 -bench=.
// PASS
// BenchmarkRun1-4  1000000              1342 ns/op
// BenchmarkRun2-4  3000000               424 ns/op
// ok      github.com/mikan/gopl/ex01/ex03 3.984s

func TestEcho3(t *testing.T) {
	run1(nil)
	run2(nil)
	run1([]string{})
	run2([]string{})
	run1(sampleInput)
	run2(sampleInput)
}
