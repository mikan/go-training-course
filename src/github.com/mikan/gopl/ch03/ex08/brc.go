// Copyright 2016 mikan. All rights reserved.

package main

import (
	"math"
	"math/big"
)

type BigRatComplex struct {
	R *big.Rat
	I *big.Rat
}

func NewBigRatComplex(x, y float64) *BigRatComplex {
	var c BigRatComplex
	c.R = new(big.Rat).SetFloat64(x)
	c.I = new(big.Rat).SetFloat64(y)
	return &c
}

func AddBigRatComplex(x, y *BigRatComplex) *BigRatComplex {
	c := NewBigRatComplex(0, 0)
	c.R.Add(x.R, y.R)
	c.I.Add(x.I, y.I)
	return c
}

func MulBigRatComplex(x, y *BigRatComplex) *BigRatComplex {
	c := NewBigRatComplex(0, 0)
	c.R.Sub(new(big.Rat).Mul(x.R, y.R), new(big.Rat).Mul(x.I, y.I))
	c.I.Add(new(big.Rat).Mul(x.R, y.I), new(big.Rat).Mul(x.I, y.R))
	return c
}

func AbsBigRatComplex(v *BigRatComplex) float64 {
	vr, _ := v.R.Float64()
	vi, _ := v.I.Float64()
	return math.Hypot(vr, vi)
}
