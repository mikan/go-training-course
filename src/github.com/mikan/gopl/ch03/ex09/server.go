// Copyright 2015-2016 mikan. All rights reserved.

// Server displays the surface.
package main

import (
	"image/png"
	"log"
	"net/http"

	"github.com/mikan/util/conv"
)

const (
	defaultX = 0.0
	defaultY = 0.0
	defaultZ = 2.0
)

// Usage: http://localhost:8080/?x=0.0&y=0.0&z=2.0
// Sample:
// http://localhost:8000/?x=-0.11&y=-0.9&z=0.01 (Top of the fractal with ultra zoom)
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		x := conv.SafeAtoF(r.FormValue("x"), defaultX)
		y := conv.SafeAtoF(r.FormValue("y"), defaultY)
		z := conv.SafeAtoF(r.FormValue("z"), defaultZ)
		w.Header().Set("Content-Type", "image/png")
		png.Encode(w, draw(x, y, z))
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
