package kata

import (
	"fmt"
)

func ArrayDiff(a, b []int) []int {
	var m = make(map[int][]int)
	var result = []int{}

	for _, el := range b {
		if _, ok := m[el]; !ok {
			m[el] = []int{}
		}

		m[el] = append(m[el], el)
	}

	for _, elb := range a {
		if _, ok := m[elb]; !ok {
			result = append(result, elb)
		}
	}
	fmt.Println(result)

	return result
}
