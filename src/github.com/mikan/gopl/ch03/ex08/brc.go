// Copyright 2016 mikan. All rights reserved.

package main

import (
	"fmt"
	"math"
	"math/big"
	"os"
)

type BigRatComplex struct {
	R *big.Rat
	I *big.Rat
}

func NewBigRatComplex(x, y float64) *BigRatComplex {
	var c BigRatComplex
	rr := new(big.Rat)
	ri := new(big.Rat)
	rr.SetFloat64(x)
	ri.SetFloat64(y)
	c.R = rr
	c.I = ri
	return &c
}

func AddBigRatComplex(x, y *BigRatComplex) *BigRatComplex {
	c := NewBigRatComplex(0, 0)
	fmt.Fprintf(os.Stderr, "[DEBUG] AddBigRatComplex: c.R.Add(%v,%v)\n", x.R, y.R)
	c.R.Add(x.R, y.R)
	fmt.Fprintf(os.Stderr, "[DEBUG] AddBigRatComplex: c.I.Add(%v,%v)\n", x.I, y.I)
	c.I.Add(x.I, y.I)
	return c
}

func MulBigRatComplex(x, y *BigRatComplex) *BigRatComplex {
	c := NewBigRatComplex(0, 0)
	r1 := new(big.Rat)
	r2 := new(big.Rat)
	fmt.Fprintf(os.Stderr, "[DEBUG] MulBigRatComplex: r1.Mul(%v,%v)\n", x.R, y.R)
	r1.Mul(x.R, y.R)
	fmt.Fprintf(os.Stderr, "[DEBUG] MulBigRatComplex: r2.Mul(%v,%v)\n", x.I, y.I)
	r2.Mul(x.I, y.I)
	fmt.Fprintf(os.Stderr, "[DEBUG] MulBigRatComplex: c.R.Sub(%v,%v)\n", r1, r2)
	c.R.Sub(r1, r2)
	i1 := new(big.Rat)
	i2 := new(big.Rat)
	fmt.Fprintf(os.Stderr, "[DEBUG] MulBigRatComplex: i1.Mul(%v,%v)\n", x.R, y.I)
	i1.Mul(x.R, y.I)
	fmt.Fprintf(os.Stderr, "[DEBUG] MulBigRatComplex: i2.Mul(%v,%v)\n", x.I, y.R)
	i2.Mul(x.I, y.R)
	fmt.Fprintf(os.Stderr, "[DEBUG] MulBigRatComplex: c.I.Add(%v,%v)\n", i1, i2)
	c.I.Add(i1, i2)
	return c
}

func AbsBigRatComplex(v *BigRatComplex) float64 {
	vr, _ := v.R.Float64()
	vi, _ := v.I.Float64()
	return math.Hypot(vr, vi)
}
