package ex06

import (
	"testing"

	orig "github.com/mikan/go-training-course/ch02/ex03/popcount"
	ex04 "github.com/mikan/go-training-course/ch02/ex04/popcount3"
	ex05 "github.com/mikan/go-training-course/ch02/ex05/popcount4"
)

func BenchmarkPopCountOrig(b *testing.B) {
	for i := 0; i < b.N; i++ {
		orig.PopCount(0x0f0f)
	}
}

func BenchmarkPopCountEx04(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ex04.PopCount(0x0f0f)
	}
}

func BenchmarkPopCountEx05(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ex05.PopCount(0x0f0f)
	}
}

/*
% go test -bench .
goos: darwin
goarch: amd64
pkg: github.com/mikan/go-training-course/ch11/ex06
cpu: Intel(R) Core(TM) m5-6Y54 CPU @ 1.10GHz
BenchmarkPopCountOrig-4         1000000000               0.3807 ns/op
BenchmarkPopCountEx04-4         45270223                25.47 ns/op
BenchmarkPopCountEx05-4         287445453                4.170 ns/op
PASS
ok      github.com/mikan/go-training-course/ch11/ex06   3.839s
 */
