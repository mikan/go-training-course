// Copyright 2016 mikan. All rights reserved.

package main

import "fmt"

var prerequisites = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	for i, course := range topologicalSort(prerequisites) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topologicalSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items map[string]bool)

	visitAll = func(items map[string]bool) {
		for item, _ := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(s2m(m[item]))
				order = append(order, item)
			}
		}
	}

	keys := make(map[string]bool)
	for key := range m {
		keys[key] = true
	}
	visitAll(keys)
	return order
}

// s2m converts slice to map
func s2m(s []string) (m map[string]bool) {
	m = make(map[string]bool)
	for _, k := range s {
		m[k] = true
	}
	return
}
