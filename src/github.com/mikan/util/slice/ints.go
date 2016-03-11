// Copyright 2016 mikan. All rights reserved.

package slice

type IntSlice []int

func (s *IntSlice) Add(x int) {
	*s = append(*s, x)
}

func (s *IntSlice) ToSlice() []int {
	return *s
}