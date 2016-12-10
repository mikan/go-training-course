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
	c.R.Sub(new(big.Float).Mul(x.R, y.R), new(big.Float).Mul(x.I, y.I))
	c.I.Add(new(big.Float).Mul(x.R, y.I), new(big.Float).Mul(x.I, y.R))
	return c
}

func AbsBigFloatComplex(v *BigFloatComplex) float64 {
	vr, _ := v.R.Float64()
	vi, _ := v.I.Float64()
	return math.Hypot(vr, vi)
}
