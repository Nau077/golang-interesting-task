// задание: переместить нули в конец
// MoveZeros([]int{1, 2, 0, 1, 0, 1, 0, 3, 0, 1}) // returns []int{ 1, 2, 1, 1, 3, 1, 0, 0, 0, 0 }

package main

import (
	"fmt"
)

func MoveZeros(arr []int) []int {
	var a = []int{}
	var b = []int{}

	for _, el := range arr {
		if el != 0 {
			a = append(a, el)
		} else {
			b = append(b, el)
		}
	}

	var res = make([]int, len(arr))

	_ = copy(res, b)
	_ = copy(res, a)

	//   res = append(res, a...)
	//   res = append(res, b...)
	fmt.Println(res)

	return res
}
