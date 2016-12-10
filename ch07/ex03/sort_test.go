// Copyright 2016 mikan. All rights reserved.

package treesort

import (
	"math/rand"
	"sort"
	"testing"
)

func TestSort(t *testing.T) {
	data := make([]int, 50)
	for i := range data {
		data[i] = rand.Int() % 50
	}
	str := Sort(data)
	if !sort.IntsAreSorted(data) {
		t.Errorf("not sorted: %v", data)
	}

	// CH7-EX3
	expected := `
10--1--0
│  └--1
└--21--20--16--15--11--10
    │  │  │  │  └--12
    │  │  │  └--15
    │  │  └--18--16
    │  │      └--19
    │  └--20
    └--37--34--24--23--22--21
        │  │  │  └--23
        │  │  └--31--28
        │  │      └--33--32
        │  └--36
        └--48--37
            └--49`

	if str != expected {
		t.Errorf("got %v want %v", str, expected)
	}
}
