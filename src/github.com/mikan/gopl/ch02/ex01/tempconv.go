// Copyright 2015-2016 mikan. All rights reserved.

// Package tempconv performs Celsius and Fahrenheit conversions.
package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string { return fmt.Sprintf("%g℃", c) }
func (f Fahrenheit) String() string { return fmt.Sprint("%g℉", f) }
