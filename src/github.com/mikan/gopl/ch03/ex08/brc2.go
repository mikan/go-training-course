// Copyright 2016 mikan. All rights reserved.

package main

import "math/big"

func AddBigRatComplex2(x, y *BigRatComplex) *BigRatComplex {
	c := NewBigRatComplex(0, 0)
	c.R.Add(x.R, y.R)
	c.I.Add(x.I, y.I)
	return reduceFloat64Space(c)
}

func MulBigRatComplex2(x, y *BigRatComplex) *BigRatComplex {
	c := NewBigRatComplex(0, 0)
	c.R.Sub(new(big.Rat).Mul(x.R, y.R), new(big.Rat).Mul(x.I, y.I))
	c.I.Add(new(big.Rat).Mul(x.R, y.I), new(big.Rat).Mul(x.I, y.R))
	return reduceFloat64Space(c)
}

func reduceFloat64Space(c *BigRatComplex) *BigRatComplex {
	fr, _ := c.R.Float64()
	c.R = new(big.Rat).SetFloat64(fr)
	fi, _ := c.I.Float64()
	c.I = new(big.Rat).SetFloat64(fi)
	return c
}
