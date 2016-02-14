// Copyright 2015-2016 mikan. All rights reserved.

// Server displays the surface.
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml")
		fmt.Fprintf(w, svg())
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
