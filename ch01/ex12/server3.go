// Copyright 2015-2016 mikan. All rights reserved.

// Server3 is a minimal "echo" and counter server.
package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var cycles = 5
		strCycles := r.FormValue("cycles")
		if len(strCycles) != 0 {
			i, err := strconv.Atoi(strCycles)
			cycles = i
			if err != nil {
				fmt.Fprintf(w, "ERROR: %s\n", err)
				return
			}
		}
		if cycles < 1 {
			fmt.Fprintf(w, "ERROR: Cycles must be specified by natrual number: %d\n", cycles)
			return
		}
		lissajous(w, cycles)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
