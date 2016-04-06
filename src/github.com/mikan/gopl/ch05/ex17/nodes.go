// Copyright 2016 mikan. All rights reserved.

package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s url\n", os.Args[0])
		return
	}
	resp, err := http.Get(os.Args[1])
	if err != nil {
		println(err)
		return
	}
	defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)
	if err != nil {
		println(err)
		return
	}
	images := ElementsByTagName(doc, "img")
	headings := ElementsByTagName(doc, "h1", "h2", "h3", "h4")
	fmt.Println("%d <img> found.", len(images))
	fmt.Println("%d <h1,h2,h3,h4> found.", len(headings))
}

func ElementsByTagName(doc *html.Node, name ...string) []*html.Node {
	if len(name) < 0 {
		return nil
	}
	return visit(nil, doc, name)
}

func visit(matches []*html.Node, n *html.Node, name []string) []*html.Node {
	if n.Type == html.ElementNode && contains(n.Data, name) {
		matches = append(matches, n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		matches = visit(matches, c, name)
	}
	return matches
}

func contains(e string, s []string) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}
