// Copyright 2016 mikan. All rights reserved.

package main

import (
	"flag"
	"fmt"

	"github.com/mikan/gopl/ch07/ex06/tempconv"
)

var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperature")
var temp2 = Celsius2Flag("temp2", 20.0, "the temperature")

func main() {
	// Description:
	// Interface "flag.Value" requires "String() string" function that used for help message. Therefore, default
	// values are displayed as value of String().
	flag.Parse()
	fmt.Println(*temp)
	fmt.Println(*temp2)
	// > ex07 -help
	// Usage of bin\ex07:
	//   -temp value
	//         the temperature (default 20°C)
	//   -temp2 value
	//         the temperature (default Celsius2#String() invoked!)
}

type Celsius2 float64
type celsius2Flag struct{ Celsius2 }

func (f *celsius2Flag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit) // no error check needed
	switch unit {
	case "C", "°C":
		f.Celsius2 = Celsius2(value)
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

func (c Celsius2) String() string { return "Celsius2#String() invoked!" }

func Celsius2Flag(name string, value Celsius2, usage string) *Celsius2 {
	f := celsius2Flag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius2
}
