// Copyright 2016 mikan. All rights reserved.

// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"math/cmplx"
)

func draw(centerX, centerY, zoom float64) *image.RGBA {
	const width, height = 1024, 1024
	var xMin, yMin, xMax, yMax = -zoom, -zoom, +zoom, +zoom
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(yMax-yMin) + yMin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xMax-xMin) + xMin
			z := complex(x+centerX, y+centerY)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	return img
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15
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
