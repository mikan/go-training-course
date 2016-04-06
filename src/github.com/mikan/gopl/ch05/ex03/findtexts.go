// Copyright 2016 mikan. All rights reserved.

// Findtexts prints the content of text nodes in an HTML document read from specified url.
package main

import (
	"fmt"
	"os"

	"bytes"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/mikan/util/input"
	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(fetch(input.Word("Input URL")))
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
		os.Exit(1)
	}
	for _, v := range visit(nil, doc) {
		fmt.Printf("%s\n", v)
	}
}

func visit(contents []string, n *html.Node) []string {
	tags := map[string]bool{"script": true, "style": true}
	if n.Type == html.TextNode && !tags[n.Parent.Data] {
		contents = append(contents, n.Data)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		contents = visit(contents, c)
	}
	return contents
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
