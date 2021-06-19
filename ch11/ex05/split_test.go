package split

import (
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	tests := []struct {
		s    string
		sep  string
		want int
	}{
		{s: "a:b:c", sep: ":", want: 3},
		{s: "", sep: ":", want: 1},
		{s: "a:b:c", sep: "", want: 5},
		{s: "a,b,c,d,e", sep: ",", want: 5},
	}
	for i, test := range tests {
		words := strings.Split(test.s, test.sep)
		if got, want := len(words), test.want; got != want {
			t.Errorf("#%d Split(%q, %q), returned %d words, want %d", i+1, test.s, test.sep, got, want)
		}
	}
}
