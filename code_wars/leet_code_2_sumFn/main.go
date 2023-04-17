package main

import "fmt"

// func twoSum(nums []int, target int) []int {
// 	m := map[int][]int{}

// 	for index, el := range nums {
// 		for i, e := range nums {

// 			if _, ok := m[el+e]; !ok {
// 				m[el+e] = []int{}
// 			}

// 			if i == index {
// 				continue
// 			}

// 			if (el + e) == target {
// 				return []int{i, index}
// 			}

// 		}
// 	}

// 	return []int{}
// }

func main() {
	nums := []int{3, 2, 4}
	target := 6
	fmt.Println(twoSum(nums, target))
}

// cупер умное решение
// func twoSum(nums []int, target int) []int {
// 	lookup := make(map[int]int)
// 	for i, v := range nums {
// 		j, ok := lookup[-v]
// 		lookup[v-target] = i
// 		if ok {
// 			return []int{j, i}
// 		}
// 	}
// 	return []int{}
// }

func twoSum(nums []int, target int) []int {
	m := map[int]int{}

	for i, el := range nums {

		x := target - el

		if j, ok := m[x]; ok {
			return []int{j, i}
		}

		m[el] = i

	}

	return []int{}
}
