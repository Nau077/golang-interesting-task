package main

import (
	"fmt"
	// "math/rand"
)

var s = []int{6, 123, 2}

func quicksort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	base := a[len(a)-1]
	var left = []int{}
	var right = []int{}

	for _, el := range a[:len(a)-1] {
		if el < base {
			left = append(left, el)
		} else {
			right = append(right, el)
		}
	}

	left = quicksort(left)
	right = quicksort(right)

	var result = []int{}
	result = append(result, left...)
	result = append(result, base)
	result = append(result, right...)

	return result
}

func main() {
	// fmt.Println(s[:len(s)-1]) вернуть массив с удалённым последним элементом
	fmt.Println(quicksort(s))
}
