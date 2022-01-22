package main

import (
	"fmt"
	"strings"
)

var prereqs = map[string][]string{
	"algorithms":     {"data structures"},
	"calculus":       {"linear algebra"},
	"linear algebra": {"calculus"},
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

func index(s string, slice []string) (int, error) {
	for i, v := range slice {
		if s == v {
			return i, nil
		}
	}

	return 0, fmt.Errorf("Not found")
}

func topoSort(m map[string][]string) (order []string, err error) {
	seen := make(map[string]bool)
	var visitAll func(items, parents []string)

	visitAll = func(items, parents []string) {
		for _, item := range items {
			r, s := seen[item]
			if s && !r {
				start, _ := index(item, parents)
				err = fmt.Errorf("Cycle: %s", strings.Join(append(parents[start:], item), " -> "))
			}
			if !s {
				seen[item] = false
				visitAll(m[item], append(parents, item))
				seen[item] = true
				order = append(order, item)
			}
		}
	}
	for key := range m {
		if err != nil {
			return nil, err
		}
		visitAll([]string{key}, nil)
	}

	return order, nil
}

func main() {
	res, err := topoSort(prereqs)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}
