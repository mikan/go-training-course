// Copyright 2016 mikan. All rights reserved.

package eval

import (
	"fmt"
	"math"
	"testing"
)

func TestEval(t *testing.T) {
	tests := []struct {
		expr string
		env  Env
		want string
	}{
		{"sqrt(A / pi)", Env{"A": 87616, "pi": math.Pi}, "167"},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 12, "y": 1}, "1729"},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 9, "y": 10}, "1729"},
		{"5 / 9 * (F - 32)", Env{"F": -40}, "-40"},
		{"5 / 9 * (F - 32)", Env{"F": 32}, "0"},
		{"5 / 9 * (F - 32)", Env{"F": 212}, "100"},
		//!-Eval
		// additional tests that don't appear in the book
		{"-1 + -x", Env{"x": 1}, "-2"},
		{"-1 - x", Env{"x": 1}, "-2"},
		//!+Eval
	}
	var prevExpr string
	for _, test := range tests {
		// Print expr only when it changes.
		if test.expr != prevExpr {
			fmt.Printf("\n%s\n", test.expr)
			prevExpr = test.expr
		}
		expr, err := Parse(test.expr)
		if err != nil {
			t.Error(err) // parse error
			continue
		}
		got := fmt.Sprintf("%.6g", expr.Eval(test.env))
		fmt.Printf("\t%s => %s\n", test.env, got)
		if got != test.want {
			t.Errorf("%s.Eval() in %s = %q, want %q\n",
				test.expr, test.env, got, test.want)
		}
	}
}

func TestErrors(t *testing.T) {
	for _, test := range []struct{ expr, wantErr string }{
		{"x % 2", "unexpected '%'"},
		{"math.Pi", "unexpected '.'"},
		{"!true", "unexpected '!'"},
		{`"hello"`, "unexpected '\"'"},
		{"log(10)", `unknown function "log"`},
		{"sqrt(1, 2)", "call to sqrt has 2 args, want 1"},
	} {
		expr, err := Parse(test.expr)
		if err == nil {
			vars := make(map[Var]bool)
			err = expr.Check(vars)
			if err == nil {
				t.Errorf("unexpected success: %s", test.expr)
				continue
			}
		}
		fmt.Printf("%-20s%v\n", test.expr, err) // (for book)
		if err.Error() != test.wantErr {
			t.Errorf("got error %s, want %s", err, test.wantErr)
		}
	}
}
