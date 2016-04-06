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

	"github.com/mikan/util/input"
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
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	first := n.FirstChild
	if first != nil {
		links = visit(links, first)
		n.RemoveChild(first)      // Remove visited child node
		n.Type = html.CommentNode // Remove current node
		links = visit(links, n)
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
