// Copyright 2015-2016 mikan. All rights reserved.

// Cf converts its numeric argument to Celsius and Fahrenheit.
package main

import (
	"fmt"
	"github.com/mikan/go-training-course/ch02/ex01/tempconv"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		k := tempconv.Kelvin(t)
		fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
		fmt.Printf("%s = %s, %s = %s\n", k, tempconv.KToC(k), c, tempconv.CToK(c))
	}
}
