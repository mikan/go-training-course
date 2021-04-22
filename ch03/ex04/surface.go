// Copyright 2015-2016 mikan. All rights reserved.

// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"log"
	"math"
)

var (
	width, height = 600, 320                     // a canvas size in pixels
	cells         = 100                          // number of grid cells
	xyRange       = 30.0                         // axis range (-xyRange..+xyRange)
	xyScale       = float64(width) / 2 / xyRange // pixels per x or y unit
	zScale        = float64(height) * 0.4        // pixels per z unit
	angle         = math.Pi / 6                  // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func svg(w, h int) string {
	width = w
	height = h
	var svg string
	svg += fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: gray; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>\n", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			if !validate(ax, ay, bx, by, cx, cy, dx, dy) {
				log.Printf("Invalid number skipped: %g,%g %g,%g %g,%g %g,%g", ax, ay, bx, by, cx, cy, dx, dy)
				continue
			}
			svg += fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g' style='fill: %s' />\n",
				ax, ay, bx, by, cx, cy, dx, dy, colorIndex(i, j))
		}
	}
	svg += "</svg>"
	return svg
}

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyRange * (float64(i)/float64(cells) - 0.5)
	y := xyRange * (float64(j)/float64(cells) - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx, sy).
	sx := float64(width/2) + (x-y)*cos30*xyScale
	sy := float64(height/2) + (x+y)*sin30*xyScale - z*zScale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

// Returns valid (all values are finite) or invalid (infinite or NaN included).
func validate(values ...float64) bool {
	for _, v := range values {
		if math.IsNaN(v) || math.IsInf(v, 0) {
			return false
		}
	}
	return true
}

// Returns color code based on height.
func colorIndex(i, j int) string {
	// z: min = -0.21722891503668823, max = 0.9850673555377986 -> (0.0 to 1.0) -> 0 to 255
	const minValue = 0.3
	const maxValue = 1.0
	z := f(xyRange*(float64(i)/float64(cells)-0.5), xyRange*(float64(j)/float64(cells)-0.5))
	r := (z + minValue) * (256 / (minValue + maxValue))
	g := 0
	b := 255 - r
	return fmt.Sprintf("#%02x%02x%02x", int(r), int(g), int(b))
}
