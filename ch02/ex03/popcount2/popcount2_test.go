// Copyright 2015-2016 mikan. All rights reserved.

package popcount2

import "testing"

func TestPopCount(t *testing.T) {
	verify(t, 0, 0)     //    0 -> 0000 0000 0000 ->  0
	verify(t, 1, 1)     //    1 -> 0000 0000 0001 ->  1
	verify(t, 2, 1)     //    2 -> 0000 0000 0010 ->  1
	verify(t, 3, 2)     //    3 -> 0000 0000 0011 ->  2
	verify(t, 255, 8)   //  255 -> 0000 1111 1111 ->  8
	verify(t, 1023, 10) // 1023 -> 0011 1111 1111 -> 10
}

func verify(t *testing.T, input uint64, expected int) {
	actual := PopCount(input)
	if expected != actual {
		t.Errorf("got %v\nwant %v (input = %d)", actual, expected, input)
	}
}
