// Copyright 2015, 2016 mikan. All rights reserved.

package tempconv

import (
	"testing"
)

func TestCToK(t *testing.T) {
	expected := Kelvin(273.15) // C -> K
	actual := CToK(Celsius(0))
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestCToF(t *testing.T) {
	expected := Fahrenheit(1)
	actual := CToF(FToC(expected)) // F -> C -> F
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestFtoK(t *testing.T) {
	expected := Kelvin(100)
	actual := FToK(CToF(KToC(expected))) // K -> C -> F -> K
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
