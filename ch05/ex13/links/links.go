// Copyright 2016 mikan. All rights reserved.

// Package links provides a link-extraction function.
package links

import (
	"fmt"
	"net/http"

	"bytes"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

	"golang.org/x/net/html"
)

var target = ""

// Extract makes an HTTP GET request to the specified URL, parses the response as HTML, and returns the links in the
// HTML document.
func Extract(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading %s: %v", url, err)
	}
	path, _, err := store(resp, bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("saving %s: %v", path, err)
	}
	doc, err := html.Parse(bytes.NewReader(data))
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue // ignore bad URLs
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return links, nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

const sep string = string(filepath.Separator)

func store(resp *http.Response, reader io.Reader) (filename string, n int64, err error) {
	if target == "" {
		target = resp.Request.Host
	} else if target != resp.Request.Host {
		return "", 0, nil
	}
	local := path.Base(resp.Request.URL.Path)
	if local == "/" || local == "." {
		local = "index.html"
	}
	path := "out" + sep + resp.Request.Host + sep + strings.Replace(local, "/", sep, -1)
	println("save path: " + path)
	err = os.MkdirAll(path[0:strings.LastIndex(path, "/")], 0777)
	if err != nil {
		return "", 0, err
	}
	f, createErr := os.Create(path)
	if createErr != nil {
		return "", 0, createErr
	}
	defer f.Close()
	n, err = io.Copy(f, reader)
	println(n)
	return path, n, err
}
