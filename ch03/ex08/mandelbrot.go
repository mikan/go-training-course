// Copyright 2016 mikan. All rights reserved.

// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"os"
	"strconv"
)

// Parameter samples...
//
// Giza-giza in C128:
// > bin/ex08 -0.100114992430011 -0.87910000000147 0.00000000000001 1 > out/m1.png && open out/m1.png
//
// Giza-giza in C64:
// > bin/ex08 -0.100114992430011 -0.87910000000147 0.000001 2 > out/m2.png && open out/m2.png
//
// BigFloat (320x320):
// > bin/ex08 -0.100114992430011 -0.87910000000147 0.000001 3 320 > out/m3.png && open out/m3.png
//
// BigRat (4x4, SUPER SLOW!!!!):
// > bin/ex08 -0.100114992430011 -0.87910000000147 0.000001 4 4 > out/m4.png && open out/m4.png
//
// BigRat with reduce:
// > bin/ex08 -0.100114992430011 -0.87910000000147 0.000001 5 320 > out/m5.png && open out/m5.png
//
func main() {
	x, y := 0.0, 0.0
	z := 2.0
	t := 1
	s := 1024
	i := 200
	switch len(os.Args[1:]) {
	case 1: // t
		t = SafeRangedAtoI(os.Args[4], t, 1, 5)
	case 2: // z t
		z = SafeAtoF(os.Args[3], z)
		t = SafeRangedAtoI(os.Args[4], t, 1, 5)
	case 4: // x y z t
		x = SafeAtoF(os.Args[1], x)
		y = SafeAtoF(os.Args[2], y)
		z = SafeAtoF(os.Args[3], z)
		t = SafeRangedAtoI(os.Args[4], t, 1, 5)
	case 5: // x y z t s
		x = SafeAtoF(os.Args[1], x)
		y = SafeAtoF(os.Args[2], y)
		z = SafeAtoF(os.Args[3], z)
		t = SafeRangedAtoI(os.Args[4], t, 1, 5)
		s = SafeAtoI(os.Args[5], s)
	case 6: // x y z t s i
		x = SafeAtoF(os.Args[1], x)
		y = SafeAtoF(os.Args[2], y)
		z = SafeAtoF(os.Args[3], z)
		t = SafeRangedAtoI(os.Args[4], t, 1, 5)
		s = SafeAtoI(os.Args[5], s)
		i = SafeAtoI(os.Args[6], s)
	default:
		log.Fatal("Usage: " + os.Args[0] + " x y z n (1=C128, 2=C64, 3=big.Float, 4=big.Rat) s i")
	}
	switch t {
	case 1:
		png.Encode(os.Stdout, draw(x, y, z, C128, s, i))
	case 2:
		png.Encode(os.Stdout, draw(x, y, z, C64, s, i))
	case 3:
		png.Encode(os.Stdout, draw(x, y, z, BigFloat, s, i))
	case 4:
		png.Encode(os.Stdout, draw(x, y, z, BigRat, s, i))
	case 5:
		png.Encode(os.Stdout, draw(x, y, z, BigRat2, s, i))
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
	BigRat2
)

var iterations uint8 = 200

func draw(centerX, centerY, zoom float64, numType NumType, size, i int) *image.RGBA {
	iterations = uint8(i)
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
			case BigRat2:
				img.Set(px, py, mandelbrotBigRat2(x+centerX, y+centerY)) // Super slow!!!!
			}
		}
		fmt.Fprintf(os.Stderr, ".")
	}
	fmt.Fprintln(os.Stderr)
	return img
}

const contrast = 15

func mandelbrot128(x, y float64) color.Color {
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

func mandelbrotBigRat2(x, y float64) color.Color {
	z := NewBigRatComplex(x, y)
	v := NewBigRatComplex(0, 0)
	for n := uint8(0); n < iterations; n++ {
		v = AddBigRatComplex2(MulBigRatComplex2(v, v), z)
		if AbsBigRatComplex(v) > 2 {
			b := 255 - contrast*n
			r := 255 - b
			g := 0
			return color.RGBA{uint8(r), uint8(g), uint8(b), 255}
		}
	}
	return color.RGBA{50, 128, 50, 255}
}
