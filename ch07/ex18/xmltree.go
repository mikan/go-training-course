// Copyright 2016 mikan. All rights reserved.

package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type Node interface{} // CharData or *Element
type CharData string
type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

func main() {
	dec := xml.NewDecoder(os.Stdin)
	var root *Element
	var stack []*Element
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmltree: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			var element Element
			element.Type = tok.Name
			element.Attr = tok.Attr
			if root == nil {
				root = &element
			} else {
				tail := stack[len(stack)-1]
				tail.Children = append(tail.Children, &element)
			}
			stack = append(stack, &element) // push
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:
			if len(stack) > 0 {
				tail := stack[len(stack)-1]
				tail.Children = (append(tail.Children, CharData(tok)))
			}
		}
	}
	printTree(root, 0)
}

func printTree(n Node, level int) {
	switch n := n.(type) {
	case *Element:
		printTabs(level)
		fmt.Printf("<%s>\n", n.Type.Local)
		for _, child := range n.Children {
			printTree(child, level+1)
		}
		printTabs(level)
		fmt.Printf("</%s>\n", n.Type.Local)
	case CharData:
		printTabs(level)
		fmt.Printf("%s\n", n)
	default:
		panic(fmt.Sprintf("Unknown type: %v\n", n))
	}
}

func printTabs(level int) {
	for i := 0; i < level; i++ {
		fmt.Print("\t")
	}
}
