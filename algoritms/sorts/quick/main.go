package main

import (
	"fmt"
	"math/rand"
	// "math/rand"
)

var s = []int{6, 123, 2, 1, 5659, 23}

func RemoveIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

func quicksort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	// base := a[len(a)-1]
	index := rand.Intn(len(a))
	base := a[index]
	var left = []int{}
	var right = []int{}
	b := RemoveIndex(a, index)

	for _, el := range b {
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
