// Copyright 2015-2016 mikan. All rights reserved.

package popcount4

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	var count = 0
	for x > 0 {
		removed := x & (x - 1)
		if x != removed {
			count++
		}
		x = removed
	}
	return count
}
