// Copyright 2015-2016 mikan. All rights reserved.

package weightconv

import (
	"math"
	"testing"
)

const acceptDelta float64 = 0.001

func TestPToKG(t *testing.T) {
	expected := Kilograms(4.5359237)
	actual := PtoKG(Pounds(10)) // P -> KG
	if math.Abs((float64)(expected-actual)) > acceptDelta {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestKGToP(t *testing.T) {
	expected := Pounds(1)
	actual := KGtoP(PtoKG(expected)) // P -> KG -> P
	if math.Abs((float64)(expected-actual)) > acceptDelta {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
