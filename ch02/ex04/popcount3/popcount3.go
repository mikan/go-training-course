// Copyright 2015-2016 mikan. All rights reserved.

package popcount3

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	var count = 0
	for i := 0; i < 64; i++ {
		tmp := x >> uint(i)
		if tmp&1 == 1 {
			count++
		}
	}
	return count
}
