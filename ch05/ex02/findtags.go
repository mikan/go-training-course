// Copyright 2016 mikan. All rights reserved.

// Findtags prints the number of tags in an HTML document read from specified url.
package main

import (
	"fmt"
	"os"

	"bytes"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/mikan/libmikan/input"
	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(fetch(input.Word("Input URL")))
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
		os.Exit(1)
	}
	for k, v := range visit(make(map[string]int), doc) {
		fmt.Printf("%-7s: %3d\n", k, v)
	}
}

func visit(types map[string]int, n *html.Node) map[string]int {
	if n.Type == html.ElementNode {
		types[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		types = visit(types, c)
	}
	return types
}

func fetch(url string) io.Reader {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
	}
	return bytes.NewReader(body)
}
