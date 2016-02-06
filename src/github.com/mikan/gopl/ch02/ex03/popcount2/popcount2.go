// Copyright 2015-2016 mikan. All rights reserved.

package popcount2

import (
	"fmt"
)

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	var count = 0
	for _, f := range fmt.Sprintf("%b", x) {
		if f == 49 { // '0'=48, '1'=49 in ASCII code
			count++
		}
	}
	return count
}
