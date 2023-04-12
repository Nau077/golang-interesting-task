package main

import "fmt"

func mergeSlices(arrs ...[]int) []int {
	length := 0
	for _, arr := range arrs {
		length += len(arr)
	}

	res := make([]int, length)
	h := res[0:length]

	for _, arr := range arrs {
		// res = append(res, arr...)
		copy(h, arr)
		c := len(arr)
		h = h[c:]
	}

	return res
}

func main() {
	a := []int{1, 2, 4, 60}
	b := []int{5, 32, 33, 45}
	c := []int{45, 344}

	fmt.Println(mergeSlices(a, b, c))
}
