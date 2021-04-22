// Copyright 2015-2016 mikan. All rights reserved.

// Server displays the surface.
package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml")
		x := safeAtoI(r.FormValue("x"), 600)
		y := safeAtoI(r.FormValue("y"), 320)
		fmt.Fprintf(w, svg(x, y))
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func safeAtoI(s string, defaultValue int) int {
	if s == "" {
		return defaultValue
	}
	ss, err := strconv.Atoi(s)
	if err != nil {
		return defaultValue
	}
	return ss
}
