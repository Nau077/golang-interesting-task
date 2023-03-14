package main

import "fmt"

func merging(a []int, b []int) []int {
	final := []int{}
	i := 0
	j := 0

	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}

	for ; i < len(a); i++ {
		final = append(final, a[i])
	}

	for ; j < len(b); j++ {
		final = append(final, b[j])
	}
	return final
}

func mergeSort(s []int) []int {
	// "тормозная система/базовый случай", когда нарезали массив минимально
	if len(s) < 2 {
		return s
	}
	mid := len(s) / 2
	left := []int{}
	right := []int{}
	// копируем в левую часть с начала до середины элементы
	left = append(left, s[0:mid]...)
	// копируем в правую часть с середины до конца элементы
	right = append(right, s[mid:len(s)]...)
	// собираем маленькие кусочки массива вместе

	return merging(mergeSort(left), mergeSort(right))
}

func main() {
	var s = []int{6, 123, 2, 1, 5659, 23}
	fmt.Println(mergeSort(s))
}
