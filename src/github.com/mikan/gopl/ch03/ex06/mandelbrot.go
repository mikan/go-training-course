// Copyright 2016 mikan. All rights reserved.

// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xMin, yMin, xMax, yMax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(yMax-yMin) + yMin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xMax-xMin) + xMin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, superSampling(img)) // NOTE: ignoring errors
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

// Super sampling for source image.
func superSampling(source image.Image) image.Image {
	img := image.NewRGBA(source.Bounds())
	for py := 0; py < img.Rect.Dy(); py++ {
		for px := 0; px < img.Rect.Dx(); px++ {
			img.Set(px, py, average(around(source, px, py)))
		}
	}
	return img
}

// Get list of around (3x3) pixels of specified pixel.
func around(source image.Image, px, py int) [9]color.Color {
	var colors [9]color.Color
	// Matrix (* = target)
	// 1 2 3
	// 4 * 6
	// 7 8 9
	colors[0] = source.At(px-1, py-1) // 1
	colors[1] = source.At(px, py-1)   // 2
	colors[2] = source.At(px+1, py-1) // 3
	colors[3] = source.At(px-1, py)   // 4
	colors[4] = source.At(px, py)     // *
	colors[5] = source.At(px+1, py)   // 6
	colors[6] = source.At(px-1, py+1) // 7
	colors[7] = source.At(px, py+1)   // 8
	colors[8] = source.At(px+1, py+1) // 9
	return colors                     // contains (0,0,0,0) when out of table.
}

func average(colors [9]color.Color) color.Color {
	var r, g, b, n uint32
	for _, c := range colors {
		cr, cg, cb, _ := c.RGBA()
		if cr == 0 && cg == 0 && cb == 0 {
			continue // skip corner case
		}
		r += cr
		g += cg
		b += cb
		n++
	}
	// see color.rgbaModel() source
	return color.RGBA{uint8(r / n >> 8), uint8(g / n >> 8), uint8(b / n >> 8), 255}
}
