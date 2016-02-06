package main
import (
	"testing"
	"github.com/mikan/gopl/ch02/ex03/popcount"
	"github.com/mikan/gopl/ch02/ex03/popcount2"
)

func TestAllPopCounts(t *testing.T) {
	p1 := popcount.PopCount(1023)
	p2 := popcount2.PopCount(1023)
	if p1 != p2 {
		t.Errorf("not equals original popcount(result = %d) and popcount2 (result = %d).", p1, p2)
	}
}

func BenchmarkPopCount1(b *testing.B) {
	loop(1, 1000000, 1023, nil)
}

func BenchmarkPopCount2(b *testing.B) {
	loop(2, 1000000, 1023, nil)
}