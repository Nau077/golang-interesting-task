package main

import (
	"fmt"
	"sort"
)

// var hostMapT = make(map[string]int)
var hostMap = map[string]int{
	"192.168.164.01": 2,
	"0.0.0.0":        5,
	"192.165.33.00":  56,
	"localhost":      1,
}

func main() {
	var hosts = []string{}

	for key := range hostMap {
		hosts = append(hosts, key)
	}

	sort.Slice(hosts, func(i, j int) bool {
		return hostMap[hosts[i]] < hostMap[hosts[j]]
	})

	fmt.Println(hosts)

	// оптимизированный вариант

	var result = [][]string{}
	for len(result) < 57 {
		result = append(result, []string{})
	}

	for key, val := range hostMap {
		result[val] = append(result[val], key)
	}

	fmt.Println(result)
}
