package main

func twoSum(nums []int, target int) []int {
	m := map[int][]int{}

	for index, el := range nums {
		for i, e := range nums {

			if _, ok := m[el+e]; !ok {
				m[el+e] = []int{}
			}

			if i == index {
				continue
			}

			if (el + e) == target {
				return []int{i, index}
			}

		}
	}

	return []int{}
}

func main() {
	nums := []int{3, 2, 4}
	target := 6
	twoSum(nums, target)
}

// cупер умное решение
// func twoSum(nums []int, target int) []int {
// 	count := make(map[int]int)
// 	for i, num := range nums {
// 		j, ok := count[num]
// 		if ok {
// 			return []int{j, i}
// 		}
// 		count[target-num] = i
// 	}
// 	return []int{}

// }
