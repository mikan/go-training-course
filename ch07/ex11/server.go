// Copyright 2016 mikan. All rights reserved.

// Http4 is an e-commerce server that registers the /list and /price
// endpoint by calling http.HandleFunc.
package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/update", db.update)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if price, ok := db[item]; ok {
		fmt.Fprintf(w, "%s\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	strNewPrice := req.URL.Query().Get("price")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	if strNewPrice == "" {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "price is empty")
		return
	}
	newPrice, err := strconv.ParseFloat(strNewPrice, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "item: %v", err)
		return
	}
	if newPrice < 0 {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "price must be 0 or larger")
		return
	}
	db[item] = dollars(newPrice) // unsafe operation in multi-thread
	fmt.Fprintf(w, "%s: %v -> %v\n", item, price, db[item])
}
