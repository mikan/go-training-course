// Copyright 2015-2016 mikan. All rights reserved.

package main

import (
	"testing"

	"github.com/mikan/go-training-course/ch02/ex03/popcount"
	"github.com/mikan/go-training-course/ch02/ex04/popcount3"
)

func TestAllPopCounts(t *testing.T) {
	p1 := popcount.PopCount(1023)
	p3 := popcount3.PopCount(1023)
	if p1 != p3 {
		t.Errorf("not equals original popcount(result = %d) and popcount3 (result = %d).", p1, p3)
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

// Result:
// >go test github.com/mikan/gopl/ch02/ex04 -bench=.
// PASS
// BenchmarkPopCount1-4    50000000                28.4 ns/op
// BenchmarkPopCount2-4    30000000                58.5 ns/op
// BenchmarkPopCount3-4     3000000               412 ns/op
// ok      github.com/mikan/gopl/ch02/ex04 5.992s
