package main

import "fmt"

func mergeSortSlice(a []int, b []int) []int {
	var N = len(a)
	var M = len(b)
	var i int
	var j int
	// c := make([]int, len(a)+len(b))
	c := []int{}

	for i < N && j < M {
		if a[i] <= b[j] {
			c = append(c, a[i])
			i += 1
		}

		if a[i] > b[j] {
			c = append(c, b[j])
			j += 1
		}
	}

	return c
}

func main() {
	a := []int{1, 2, 4, 60}
	b := []int{5, 32, 33, 45}
	fmt.Println(mergeSortSlice(a, b))
}
