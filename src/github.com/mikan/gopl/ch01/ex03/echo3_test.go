// Copyright 2015-2016 mikan. All rights reserved.

package main

import "testing"

const loopCount = 10000000

var sampleInput = []string{"foo", "bar", "baz"}

func BenchmarkRun1(b *testing.B) {
	for i := 1; i < loopCount; i++ {
		run1(sampleInput)
	}
}

func BenchmarkRun2(b *testing.B) {
	for i := 1; i < loopCount; i++ {
		run2(sampleInput)
	}
}

// Result:
// go test github.com/mikan/gopl/ch01/ex03 -bench=.
// PASS
// BenchmarkRun1-4        1        2279895668 ns/op
// BenchmarkRun2-4        1        1707501409 ns/op
// ok      github.com/mikan/gopl/ch01/ex03 4.000s

func TestEcho3(t *testing.T) {
	run1(nil)
	run2(nil)
	run1([]string{})
	run2([]string{})
	run1(sampleInput)
	run2(sampleInput)
}
