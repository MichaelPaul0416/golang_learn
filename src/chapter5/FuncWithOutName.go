package chapter5

import "sort"

var Prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"database":              {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func TopoSort(m map[string][]string) []string{
	var order []string
	seen := make(map[string]bool)//list

	//匿名函数
	var visitAll func(k []string)

	visitAll = func(items []string) {
		for _,item := range items{
			if !seen[item]{
				seen[item] = true
				visitAll(m[item])
				order = append(order,item)
			}
		}
	}

	var keys []string
	for key := range m{
		keys = append(keys,key)
	}
	sort.Strings(keys)
	visitAll(keys)
	return order
}