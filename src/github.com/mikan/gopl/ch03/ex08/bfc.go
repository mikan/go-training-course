// Copyright 2016 mikan. All rights reserved.
// Copyright 2016 mikan. All rights reserved.

package main

import (
	"math"
	"math/big"
)

type BigFloatComplex struct {
	R *big.Float
	I *big.Float
}

// Memo http://w3e.kanazawa-it.ac.jp/math/category/fukusosuu/fukusosuu-no-sisokuenzan.html
// Memo http://w3e.kanazawa-it.ac.jp/math/category/fukusosuu/fukuso-zettaichi.html
func NewBigFloatComplex(x, y float64) *BigFloatComplex {
	var c BigFloatComplex
	c.R = big.NewFloat(x)
	c.I = big.NewFloat(y)
	return &c
}

func AddBigFloatComplex(x, y *BigFloatComplex) *BigFloatComplex {
	c := NewBigFloatComplex(0, 0)
	c.R.Add(x.R, y.R)
	c.I.Add(x.I, y.I)
	return c
}

func MulBigFloatComplex(x, y *BigFloatComplex) *BigFloatComplex {
	c := NewBigFloatComplex(0, 0)
	r1 := new(big.Float)
	r2 := new(big.Float)
	r1.Mul(x.R, y.R)
	r2.Mul(x.I, y.I)
	c.R.Sub(r1, r2)
	i1 := new(big.Float)
	i2 := new(big.Float)
	i1.Mul(x.R, y.I)
	i2.Mul(x.I, y.R)
	c.I.Add(i1, i2)
	return c
}

func AbsBigFloatComplex(v *BigFloatComplex) float64 {
	vr, _ := v.R.Float64()
	vi, _ := v.I.Float64()
	return math.Hypot(vr, vi)
}
