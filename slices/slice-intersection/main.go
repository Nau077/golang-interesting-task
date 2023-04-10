package main

import "fmt"

func intersect(a, b []int) []int {
	c := make(map[int]int)
	result := []int{}

	for _, el := range a {
		if _, ok := c[el]; ok {
			c[el] += 1
		} else {
			c[el] = 1
		}
	}

	for _, el := range b {
		if _, ok := c[el]; ok {
			result = append(result, el)
		}
	}

	return result
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := []int{3, 17, 6}

	result := intersect(a, b)
	fmt.Print(result)
}
