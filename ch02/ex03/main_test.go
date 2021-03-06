// Copyright 2015-2016 mikan. All rights reserved.

package main

import (
	"testing"

	"github.com/mikan/go-training-course/ch02/ex03/popcount"
	"github.com/mikan/go-training-course/ch02/ex03/popcount2"
)

func TestAllPopCounts(t *testing.T) {
	p1 := popcount.PopCount(1023)
	p2 := popcount2.PopCount(1023)
	if p1 != p2 {
		t.Errorf("not equals original popcount(result = %d) and popcount2 (result = %d).", p1, p2)
	}
}

func BenchmarkPopCount1(b *testing.B) {
	loop(1, b.N, 1023, nil)
}

func BenchmarkPopCount2(b *testing.B) {
	loop(2, b.N, 1023, nil)
}

// Result:
// >go test github.com/mikan/gopl/ch02/ex03 -bench=.
// PASS
// BenchmarkPopCount1-4    50000000                30.9 ns/op
// BenchmarkPopCount2-4    30000000                45.2 ns/op
// ok      github.com/mikan/gopl/ch02/ex03 3.635s
