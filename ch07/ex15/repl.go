// Copyright 2016 mikan. All rights reserved.

package main

import (
	"fmt"
	"math"
	"os"

	"strconv"

	"github.com/mikan/go-training-course/ch07/ex13/eval"
	"github.com/mikan/libmikan/input"
)

func main() {
	env := eval.Env{"pi": math.Pi}
	for {
		text := input.SingleLine("expr")
		if input.IsQuit(text) {
			break
		}
		expr, err := eval.Parse(text)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			continue
		}
		vars := make(map[eval.Var]bool)
		check := expr.Check(vars)
		if check != nil {
			fmt.Printf("%v\n", check)
			continue
		}
		for v, _ := range vars {
			found := false
			for e, _ := range env {
				if v == e {
					found = true
					break // found
				}
			}
			if !found { // not found
				for {
					strAdd := input.SingleLine(fmt.Sprintf("[Var] %s", string(v)))
					if input.IsQuit(strAdd) {
						return
					}
					add, parseError := strconv.ParseFloat(strAdd, 64)
					if parseError == nil {
						env[v] = add
						break
					}
					fmt.Fprintf(os.Stderr, "ERROR: %s\n", parseError) // continue
				}
			}
		}
		got := fmt.Sprintf("%.6g", expr.Eval(env))
		fmt.Printf("%s\n", got)
	}
}

// Example:
//
// expr > A+B
// [Var] A > 1
// [Var] B > 2
// 3
// expr > A+B+1
// 4
// expr > X+A
// [Var] X > 100
// 101
// expr > quit
