// Copyright 2016 mikan. All rights reserved.

package main

import (
	"fmt"
	"math"
	"os"

	"github.com/mikan/gopl/ch07/ex13/eval"
	"github.com/mikan/libmikan/input"
)

func main() {
	env := eval.Env{"pi": math.Pi}
	for {
		original := input.SingleLine("expr")
		if input.IsQuit(original) {
			break
		}
		originalExpr, err := eval.Parse(original)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			continue
		}
		originalGot := fmt.Sprintf("%.6g", originalExpr.Eval(env))
		format := originalExpr.String()
		formatExpr, err := eval.Parse(format)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			continue
		}
		formatGot := fmt.Sprintf("%.6g", formatExpr.Eval(env))
		fmt.Printf("original:\t%s = %s\n", original, originalGot)
		fmt.Printf("format:\t\t%s = %s\n", format, formatGot)
	}
}

// Example:
//
// expr > 1+2+3
// original:       1+2+3 = 6
// format:         ((1 + 2) + 3) = 6
// expr > sqrt(2*2)
// original:       sqrt(2*2) = 2
// format:         sqrt((2 * 2)) = 2
// expr > quit
