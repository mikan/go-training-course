// Copyright 2016 mikan. All rights reserved.

// The sum program demonstrates a variadic function.
package main

import (
	"fmt"
	"math"
)

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func max(vals ...int) int {
	if len(vals) < 0 {
		panic("max: vals is empty")
	}
	tmpMax := math.MinInt8
	for _, val := range vals {
		if val > tmpMax {
			tmpMax = val
		}
	}
	return tmpMax
}

func min(vals ...int) int {
	if len(vals) < 0 {
		panic("min: vals is empty")
	}
	tmpMin := math.MaxInt8
	for _, val := range vals {
		if val < tmpMin {
			tmpMin = val
		}
	}
	return tmpMin
}

func max2(first int, vals ...int) int {
	tmpMax := first
	for _, val := range vals {
		if val > tmpMax {
			tmpMax = val
		}
	}
	return tmpMax
}

func min2(first int, vals ...int) int {
	tmpMin := first
	for _, val := range vals {
		if val < tmpMin {
			tmpMin = val
		}
	}
	return tmpMin
}

func main() {
	fmt.Println(sum())           //  "0"
	fmt.Println(sum(3))          //  "3"
	fmt.Println(sum(1, 2, 3, 4)) //  "10"

	values := []int{1, 2, 3, 4}
	fmt.Println(sum(values...)) // "10"
}
