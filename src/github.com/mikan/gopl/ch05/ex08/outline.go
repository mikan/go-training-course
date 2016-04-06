// Copyright 2016 mikan. All rights reserved.

// Outline prints the outline of an HTML document tree.
package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}
	forEachNode(doc, startElement, endElement)
	return nil
}

// forEachNode calls the functions pre(x) and post(x) for each node x in the tree rooted at n. Both functions are
// optional.
// pre is called before the children are visited (preorder) and post is called after (postorder).
func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) bool {
	if pre != nil {
		if !pre(n) {
			return false
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if !forEachNode(c, pre, post) {
			return false
		}
	}
	if post != nil {
		if !post(n) {
			return false
		}
	}
	return true
}

var depth int

func startElement(n *html.Node) bool {
	if n.Type == html.ElementNode {
		var attrs string
		for _, attr := range n.Attr {
			attrs += fmt.Sprintf(" %s='%s'", attr.Key, attr.Val)
		}
		if n.FirstChild == nil {
			fmt.Printf("%*s<%s%s/>\n", depth*2, "", n.Data, attrs)
			return false
		} else {
			fmt.Printf("%*s<%s%s>\n", depth*2, "", n.Data, attrs)
			depth++
		}
	}
	return true
}

func endElement(n *html.Node) bool {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		if n.Data == "html" {
			return false
		}
	}
	return true
}

func ElementByID(doc *html.Node, id string) *html.Node {
	var node *html.Node
	forEachNode(doc, func(n *html.Node) bool {
		if n.Type == html.ElementNode {
			for _, attr := range n.Attr {
				if attr.Key == "id" && attr.Val == id {
					node = n
					return false
				}
			}
		}
		return startElement(n)
	}, endElement)
	return node
}
