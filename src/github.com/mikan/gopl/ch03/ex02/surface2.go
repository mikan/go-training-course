// Copyright 2015-2016 mikan. All rights reserved.

// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"log"
	"math"
	"os"
)

const (
	width, height = 600, 320            // acnvas size in pixels
	cells         = 100                 // number of grid cells
	xyRange       = 30.0                // axis range (-xyRange..+xyRange)
	xyScale       = width / 2 / xyRange // pixels per x or y unit
	zScale        = height * 0.4        // pxels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	method := "eggbox"
	if len(os.Args) == 2 {
		method = os.Args[1]
	}
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: gray; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>\n", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j, method)
			bx, by := corner(i, j, method)
			cx, cy := corner(i, j+1, method)
			dx, dy := corner(i+1, j+1, method)
			if math.IsNaN(ax) || math.IsNaN(ay) || math.IsNaN(bx) || math.IsNaN(by) ||
				math.IsNaN(cx) || math.IsNaN(cy) || math.IsNaN(dx) || math.IsNaN(dy) {
				log.Printf("NaN skipped: %g,%g %g,%g %g,%g %g,%g", ax, ay, bx, by, cx, cy, dx, dy)
				continue
			}
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n", ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int, method string) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyRange * (float64(i)/cells - 0.5)
	y := xyRange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	var z float64
	switch method {
	case "eggbox":
		z = eggBox(x, y)
	case "moguls":
		z = moguls(x, y)
	case "saddle":
		z = saddle(x, y)
	default:
		log.Fatal("Unknown method: " + method)
	}

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx, sy).
	sx := width/2 + (x-y)*cos30*xyScale
	sy := height/2 + (x+y)*sin30*xyScale - z*zScale
	return sx, sy
}

func eggBox(x, y float64) float64 {
	// EGG-BOX
	// [Cos[x]Sin[y], {x, 0, 4 Pi}, [y, 0, 4 Pi]]
	// Refer http://archive.vector.org.uk/trad/v224/v224-IAN/clark.htm
	return math.Cos(x) * math.Sin(y) / 5
}

func moguls(x, y float64) float64 {
	// Hmm... Is this moguls? (refer fig.7.7)
	return math.Pow(2, math.Sin(y)) * math.Pow(2, math.Sin(x)) / 12
}

func saddle(x, y float64) float64 {
	// Saddle
	// [x^2-y^2, {x, -2, 2}, {y, -2, 2}]
	// Refer http://www.math.harvard.edu/archive/21a_summer_03/labs/labhtml/
	return (math.Pow(x, 2)/100 - math.Pow(y, 2)/100) / 5
}
