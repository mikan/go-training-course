// Copyright 2015-2016 mikan. All rights reserved.

// Server displays the surface.
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mikan/util/conv"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml")
		x := conv.SafeAtoF(r.FormValue("x"), 600)
		y := conv.SafeAtoF(r.FormValue("y"), 320)
		fmt.Fprintf(w, svg(x, y))
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
