// Copyright 2016 mikan. All rights reserved.

package main

import (
	"strings"
	"testing"

	"golang.org/x/net/html"
)

const expectedID = "foo"

func TestElementByID(t *testing.T) {
	src := `<!DOCTYPE html>
	<html>
	<head>
	<title>Foo!</title>
	</head>
	<body>
	<p id="foo">Foo!</p>
	<p id="foo">Bar!</p>
	<p id="foo">Baz!</p>
	</body>
	</html>`
	doc, err := html.Parse(strings.NewReader(src))
	if err != nil {
		t.Errorf("error: %v", err)
		return
	}
	node := ElementByID(doc, expectedID)
	found := false
	for _, attr := range node.Attr {
		if attr.Key == "id" && attr.Val == expectedID {
			found = true
			return
		}
	}
	if !found {
		t.Errorf("id %s not found in the result", expectedID)
	}
}
