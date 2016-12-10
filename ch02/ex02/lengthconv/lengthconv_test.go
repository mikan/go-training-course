// Copyright 2015-2016 mikan. All rights reserved.

package lengthconv

import (
	"math"
	"testing"
)

const acceptDelta float64 = 0.001

func TestMToF(t *testing.T) {
	expected := Feat(32.808)
	actual := MtoF(Meters(10)) // F -> M
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestFToM(t *testing.T) {
	expected := Meters(1)
	actual := FtoM(MtoF(expected))                          // M -> F -> M
	if math.Abs((float64)(expected-actual)) > acceptDelta { // I want to "assertEquals(expected, actual, delta)"
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
