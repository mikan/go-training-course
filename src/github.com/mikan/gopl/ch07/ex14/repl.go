// Copyright 2016 mikan. All rights reserved.

package main

import (
	"fmt"
	"math"
	"os"

	"github.com/mikan/gopl/ch07/ex14/eval"
	"github.com/mikan/libmikan/input"
)

func main() {
	env := eval.Env{"pi": math.Pi}
	for {
		text := input.SingleLine("expr")
		if input.IsQuit(text) {
			break
		}
		originalExpr, err := eval.Parse(text)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			continue
		}
		got := fmt.Sprintf("%.6g", originalExpr.Eval(env))
		fmt.Printf("%s\n", got)
	}
}

// Example:
//
// expr > min(1,2)
// 1
// expr > min(2,1)
// 1
// expr > min(-100,100)
// -100
// expr > quit
