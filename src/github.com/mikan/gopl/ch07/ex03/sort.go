// Copyright 2016 mikan. All rights reserved.

// Package treesort provides insertion sort using an unbalanced binary tree.
package treesort

import (
	"fmt"
)

//!+
type tree struct {
	value       int
	left, right *tree
}

// Sort sorts values in place.
func Sort(values []int) string {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	root.String()
	appendValues(values[:0], root)
	return root.String() // CH7-EX3
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

var LEFT = 0
var RIGHT = 1
var DONE = 2

func (t *tree) String() string {
	text := "\n"
	var nodes []*tree
	var dirs []int
	nodes = append(nodes, t)
	dirs = append(dirs, LEFT)
	text += fmt.Sprintf("%d", t.value)
	for len(nodes) > 0 {
		n := nodes[len(nodes)-1]
		d := dirs[len(dirs)-1]
		if d == LEFT && n.left != nil {
			text += fmt.Sprintf("--%d", n.left.value)
			dirs[len(dirs)-1] = RIGHT
			nodes = append(nodes, n.left)
			dirs = append(dirs, LEFT)
			continue
		} else if d == RIGHT && n.right != nil {
			text += "\n"
			for j := 0; j < len(nodes); j++ {
				if j == len(nodes)-1 {
					text += "└--"
				} else if dirs[j] != DONE && nodes[j].right != nil {
					text += "│  "
				} else {
					text += "    "
				}
			}
			text += fmt.Sprintf("%d", n.right.value)
			dirs[len(dirs)-1] = DONE
			nodes = append(nodes, n.right)
			dirs = append(dirs, LEFT)
		} else {
			nodes = nodes[:len(nodes)-1]
			dirs = dirs[:len(dirs)-1]
		}
	}
	return text
}
