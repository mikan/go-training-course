// Copyright 2016 mikan. All rights reserved.

package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	src := `<!DOCTYPE html>
	<html>
	<head>
	<title>Foo!</title>
	</head>
	<body>
	<a href="a.html">Foo</a>
	<a href="b.html">Bar</a>
	<a href="c.html">Baz</a>
	</body>
	</html>`
	reader := NewReader(src)
	doc, err := html.Parse(reader)
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
