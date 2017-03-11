// Copyright 2016 mikan. All rights reserved.

// Findlinks prints the links in an HTML document read from specified url.
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
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func visit(links []string, n *html.Node) []string {
	tags := map[string]bool{"a": true, "link": true, "img": true, "script": true}
	attrs := map[string]bool{"href": true, "src": true}
	if n.Type == html.ElementNode && tags[n.Data] {
		for _, attr := range n.Attr {
			if attrs[attr.Key] {
				links = append(links, fmt.Sprintf("[%s] %s", n.Data, attr.Val))
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
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
