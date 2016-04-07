// Copyright 2016 mikan. All rights reserved.

package main

import (
	"bufio"
	"fmt"
	"net/http"
	"strings"

	"github.com/mikan/util/input"
	"golang.org/x/net/html"
)

func main() {
	words, images, err := CountWordsAndImages(input.Word("Input URL"))
	if err != nil {
		fmt.Errorf("parseing HTML: %s", err)
		return
	}
	fmt.Printf("words:  %d\n", words)
	fmt.Printf("iamges: %d\n", images)
}

// CountWordsAndImages does an HTTP GET request for the HTML document url and returns the numb er of words and images
// in it.
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err == nil {
		words, images = countWordsAndImages(doc)
	}
	return words, images, err
}

func countWordsAndImages(n *html.Node) (words, images int) {
	w, i, _ := visit(0, 0, n)
	return w, i
}

func visit(words, images int, n *html.Node) (w, i int, node *html.Node) {
	tags := map[string]bool{"script": true, "style": true}
	if n.Type == html.TextNode && !tags[n.Parent.Data] {
		words += countWords(n.Data)
	}
	if n.Type == html.ElementNode && n.Data == "img" {
		for _, tag := range n.Attr {
			if tag.Key == "src" {
				images++
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		words, images, n = visit(words, images, c)
	}
	return words, images, n
}

func countWords(s string) int {
	input := bufio.NewScanner(strings.NewReader(s))
	input.Split(bufio.ScanWords)
	var counts int
	for input.Scan() {
		input.Text()
		counts++
	}
	return counts
}
