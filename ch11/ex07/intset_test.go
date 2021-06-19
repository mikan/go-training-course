// Copyright 2021 mikan. All rights reserved.

package intset

import (
	"math"
	"math/rand"
	"testing"
	"time"
)

func BenchmarkIntSet_Add(b *testing.B) {
	seed := time.Now().UTC().UnixNano()
	//b.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < b.N; i++ {
		var set IntSet
		set.Add(rng.Intn(math.MaxInt16))
	}
}

func BenchmarkStdSet_Add(b *testing.B) {
	seed := time.Now().UTC().UnixNano()
	//b.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < b.N; i++ {
		var set StdSet
		set.Add(rng.Intn(math.MaxInt16))
	}
}

func BenchmarkIntSet_UnionWith(b *testing.B) {
	seed := time.Now().UTC().UnixNano()
	//b.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < b.N; i++ {
		var set1, set2 IntSet
		set1.Add(rng.Intn(math.MaxInt16))
		set2.Add(rng.Intn(math.MaxInt16))
		set1.UnionWith(&set2)
	}
}

func BenchmarkStdSet_UnionWith(b *testing.B) {
	seed := time.Now().UTC().UnixNano()
	//b.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < b.N; i++ {
		var set1, set2 StdSet
		set1.Add(rng.Intn(math.MaxInt16))
		set2.Add(rng.Intn(math.MaxInt16))
		set1.UnionWith(&set2)
	}
}

/*
Oops! IntSet is slower than standard map.
Result:

% go test -bench .
goos: darwin
goarch: amd64
pkg: github.com/mikan/go-training-course/ch11/ex07
cpu: Intel(R) Core(TM) m5-6Y54 CPU @ 1.10GHz
BenchmarkIntSet_Add-4             499368              2364 ns/op
BenchmarkStdSet_Add-4            8570016               131.1 ns/op
BenchmarkIntSet_UnionWith-4       197502              5571 ns/op
BenchmarkStdSet_UnionWith-4      2710684               443.7 ns/op
PASS
ok      github.com/mikan/go-training-course/ch11/ex07   6.557s
*/
