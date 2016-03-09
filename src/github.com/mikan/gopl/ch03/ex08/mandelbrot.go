// Copyright 2016 mikan. All rights reserved.

// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"math/big"
	"math/cmplx"
	"os"
	"strconv"

	"github.com/mikan/util/conv"
)

// Parameter samples...
//
// Giza-giza in C128:
// > bin/ex08 -0.100114992430011 -0.87910000000147 0.00000000000001 1 > out/m1.png && open out/m1.png
//
// Giza-giza in C64:
// > bin/ex08 -0.100114992430011 -0.87910000000147 0.00001 2 > out/m2.png && open out/m2.png
//
// BigFloat (100x100):
// > bin/ex08 -0.100114992430011 -0.87910000000147 0.00001 3 100 > out/m3.png && open out/m3.png
//
// BigRat (SUPER SLOW!!!!):
// > bin/ex08 -0.100114992430011 -0.87910000000147 0.00001 4 10 > out/m4.png && open out/m4.png
func main() {
	x, y := 0.0, 0.0
	z := 2.0
	t := 1
	s := 1024
	switch len(os.Args[1:]) {
	case 1: // t
		t = conv.SafeRangedAtoI(os.Args[4], t, 1, 4)
	case 2: // z t
		z = conv.SafeAtoF(os.Args[3], z)
		t = conv.SafeRangedAtoI(os.Args[4], t, 1, 4)
	case 4: // x y z t
		x = conv.SafeAtoF(os.Args[1], x)
		y = conv.SafeAtoF(os.Args[2], y)
		z = conv.SafeAtoF(os.Args[3], z)
		t = conv.SafeRangedAtoI(os.Args[4], t, 1, 4)
	case 5: // x y z t s
		x = conv.SafeAtoF(os.Args[1], x)
		y = conv.SafeAtoF(os.Args[2], y)
		z = conv.SafeAtoF(os.Args[3], z)
		t = conv.SafeRangedAtoI(os.Args[4], t, 1, 4)
		s = conv.SafeAtoI(os.Args[5], s)
	default:
		log.Fatal("Usage: " + os.Args[0] + " x y z n (1=C128, 2=C64, 3=big.Float, 4=big.Rat")
	}
	switch t {
	case 1:
		png.Encode(os.Stdout, draw(x, y, z, C128, s))
	case 2:
		png.Encode(os.Stdout, draw(x, y, z, C64, s))
	case 3:
		png.Encode(os.Stdout, draw(x, y, z, BigFloat, s))
	case 4:
		png.Encode(os.Stdout, draw(x, y, z, BigRat, s))
	default:
		log.Fatal("Unexpected t: " + strconv.Itoa(1))
	}
}

type NumType uint8

const (
	C64 NumType = iota
	C128
	BigFloat
	BigRat
)

func draw(centerX, centerY, zoom float64, numType NumType, size int) *image.RGBA {
	var width, height = size, size
	var xMin, yMin, xMax, yMax = -zoom, -zoom, +zoom, +zoom
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/float64(height)*(yMax-yMin) + yMin
		for px := 0; px < width; px++ {
			x := float64(px)/float64(width)*(xMax-xMin) + xMin
			switch numType {
			case C128:
				fallthrough
			default:
				img.Set(px, py, mandelbrot128(x+centerX, y+centerY))
			case C64:
				img.Set(px, py, mandelbrot64(x+centerX, y+centerY))
			case BigFloat:
				img.Set(px, py, mandelbrotBigFloat(x+centerX, y+centerY))
			case BigRat:
				img.Set(px, py, mandelbrotBigRat(x+centerX, y+centerY)) // Super slow!!!!
			}
		}
	}
	return img
}

func mandelbrot128(x, y float64) color.Color {
	const iterations = 200
	const contrast = 15
	z := complex(x, y)
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			b := 255 - contrast*n
			r := 255 - b
			g := 0
			return color.RGBA{uint8(r), uint8(g), uint8(b), 255}
		}
	}
	return color.RGBA{50, 128, 50, 255}
}

func mandelbrot64(x, y float64) color.Color {
	const iterations = 200
	const contrast = 15
	z := complex64(complex(x, y))
	var v complex64
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(complex128(v)) > 2 {
			b := 255 - contrast*n
			r := 255 - b
			g := 0
			return color.RGBA{uint8(r), uint8(g), uint8(b), 255}
		}
	}
	return color.RGBA{50, 128, 50, 255}
}

func mandelbrotBigFloat(x, y float64) color.Color {
	const iterations = 200
	const contrast = 15
	z := NewBigFloatComplex(x, y)
	v := NewBigFloatComplex(0, 0)
	for n := uint8(0); n < iterations; n++ {
		v = AddBigFloatComplex(MulBigFloatComplex(v, v), z)
		if AbsBigFloatComplex(v) > 2 {
			b := 255 - contrast*n
			r := 255 - b
			g := 0
			return color.RGBA{uint8(r), uint8(g), uint8(b), 255}
		}
	}
	return color.RGBA{50, 128, 50, 255}
}

func mandelbrotBigRat(x, y float64) color.Color {
	const iterations = 200
	const contrast = 15
	z := NewBigRatComplex(x, y)
	v := NewBigRatComplex(0, 0)
	for n := uint8(0); n < iterations; n++ {
		v = AddBigRatComplex(MulBigRatComplex(v, v), z)
		if AbsBigRatComplex(v) > 2 {
			b := 255 - contrast*n
			r := 255 - b
			g := 0
			return color.RGBA{uint8(r), uint8(g), uint8(b), 255}
		}
	}
	return color.RGBA{50, 128, 50, 255}
}

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

type BigRatComplex struct {
	R *big.Rat
	I *big.Rat
}

func AddBigRatComplex(x, y *BigRatComplex) *BigRatComplex {
	c := NewBigRatComplex(0, 0)
	c.R.Add(x.R, y.R)
	c.I.Add(x.I, y.I)
	return c
}

func MulBigRatComplex(x, y *BigRatComplex) *BigRatComplex {
	c := NewBigRatComplex(0, 0)
	r1 := new(big.Rat)
	r2 := new(big.Rat)
	r1.Mul(x.R, y.R)
	r2.Mul(x.I, y.I)
	c.R.Sub(r1, r2)
	i1 := new(big.Rat)
	i2 := new(big.Rat)
	i1.Mul(x.R, y.I)
	i2.Mul(x.I, y.R)
	c.I.Add(i1, i2)
	return c
}

func AbsBigRatComplex(v *BigRatComplex) float64 {
	vr, _ := v.R.Float64()
	vi, _ := v.I.Float64()
	return math.Hypot(vr, vi)
}
