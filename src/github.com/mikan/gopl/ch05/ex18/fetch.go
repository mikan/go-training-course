// Copyright 2016 mikan. All rights reserved.

// Fetch saves the contents of a URL into a local file.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

// Fetch downloads the URL and returns the name and length of the local file.
func fetch(url, target string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(target + string(filepath.Separator) + local)
	if err != nil {
		return "", 0, err
	}
	defer f.Close()
	n, err = io.Copy(f, resp.Body)
	return local, n, err
}

func main() {
	for _, url := range os.Args[1:] {
		local, n, err := fetch(url, "out")
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch %s: %v\n", url, err)
			continue
		}
		fmt.Fprintf(os.Stderr, "%s => %s (%d bytes).\n", url, local, n)
	}
}
