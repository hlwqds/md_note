package main

import (
	"fmt"
	"sort"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra'"},
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
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visit func(items []string)
	visit = func(items []string) {
		for _, item := range items {
			if seen[item] == false {
				seen[item] = true
				visit(m[item])
				order = append(order, item)
			}
		}
	}
	var keys []string
	for key, _ := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visit(keys)
	return order
}
