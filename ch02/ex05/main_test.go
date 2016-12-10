// Copyright 2015-2016 mikan. All rights reserved.

package main

import (
	"testing"

	"github.com/mikan/go-training-course/ch02/ex05/popcount4"
	"github.com/mikan/go-training-course/ch02/ex03/popcount"
)

func TestAllPopCounts(t *testing.T) {
	p1 := popcount.PopCount(1023)
	p4 := popcount4.PopCount(1023)
	if p1 != p4 {
		t.Errorf("not equals original popcount(result = %d) and popcount4 (result = %d).", p1, p4)
	}
}

func BenchmarkPopCount1(b *testing.B) {
	loop(1, b.N, 1023, nil)
}

func BenchmarkPopCount2(b *testing.B) {
	loop(2, b.N, 1023, nil)
}

func BenchmarkPopCount3(b *testing.B) {
	loop(3, b.N, 1023, nil)
}

func BenchmarkPopCount4(b *testing.B) {
	loop(4, b.N, 1023, nil)
}

// Result:
// >go test github.com/mikan/gopl/ch02/ex05 -bench=.
// PASS
// BenchmarkPopCount1-4    50000000                28.5 ns/op
// BenchmarkPopCount2-4    30000000                59.5 ns/op
// BenchmarkPopCount3-4     3000000               398 ns/op
// BenchmarkPopCount4-4    20000000                85.3 ns/op
// ok      github.com/mikan/gopl/ch02/ex05 7.551s
