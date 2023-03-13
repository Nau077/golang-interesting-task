package main

import (
	"fmt"
	"math/rand"
)

var s = []int{6, 123, 2, 1, 5659, 23}

func quicksort(s []int) []int {
	if len(s) < 2 {
		return s
	}
	base := rand.Intn(len(s))
	var left = []int{}
	var right = []int{}

	for _, el := range s {
		if el < s[base] {
			left = append(left, el)
		}

		if el > s[base] {
			right = append(right, el)
		}
	}

	left = quicksort(left)
	right = quicksort(right)

	result := []int{}

	result = append(result, left...)
	result = append(result, s[base])
	result = append(result, right...)

	return result

}

func main() {
	fmt.Println(quicksort(s))
}
