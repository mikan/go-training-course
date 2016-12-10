// Copyright 2016 mikan. All rights reserved.

package eval

import (
	"fmt"
	"math"
	"testing"
)

func TestString(t *testing.T) {
	patterns := []string{"1+2+3", "sqrt(2*2)"}
	env := Env{"pi": math.Pi}
	for _, p := range patterns {
		expr, err := Parse(p)
		if err != nil {
			t.Error(err)
			return
		}
		got := fmt.Sprintf("%.6g", expr.Eval(env))
		expr2, err := Parse(expr.String())
		if err != nil {
			t.Error(err)
			return
		}
		got2 := fmt.Sprintf("%.6g", expr2.Eval(env))
		if got != got2 {
			t.Errorf("got %v want %v", got, got2)
		}
	}
}
