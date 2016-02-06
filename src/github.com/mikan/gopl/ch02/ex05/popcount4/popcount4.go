// Copyright 2015-2016 mikan. All rights reserved.

package popcount4

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	var count = 0
	var tmp uint64 = x
	for tmp > 0 {
		removed := tmp & (tmp - 1)
		if tmp != removed {
			count++
		}
		tmp = removed
	}
	return count
}
