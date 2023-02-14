package main

import (
	"fmt"
	"regexp"
	"strings"
)

func indexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}

	return -1 //not found.
}

func High(s string) string {
	var arr = []string{"a", "b", "v", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m",
		"n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	r := strings.Split(s, "")

	m := map[int]int{}
	mStr := map[int]string{}
	n := 1
	m[n] = 0

	for _, el := range r {
		whitespace := regexp.MustCompile(`\s`).MatchString(string(el))

		if whitespace == true {
			n = n + 1
			m[n] = 0
		}
		vs, ok := mStr[n]

		if ok {
			mStr[n] = vs + el
		} else {
			mStr[n] = el
		}

		k := indexOf(string(el), arr)
		val, ok := m[n]

		if ok {
			if k != -1 {
				m[n] = val + k + 1
			}
		}
	}

	sum := []int{}
	// rvM := reverseMap(m)

	for _, val := range m {
		sum = append(sum, val)
	}

	arrM := findMax(m)
	// max := arrM[0]
	key := arrM[1]
	fmt.Println(m)
	fmt.Println(mStr)
	// response := strings.Join(strings.Fields(mStr[rvM[max]]), "")

	response := strings.Join(strings.Fields(mStr[key]), "")
	fmt.Println(response)
	// fmt.Println(m)
	// fmt.Println(max)

	return response
}

func findMax(m map[int]int) []int {
	var largerNumber, temp int
	var k, prevkey int

	arr := []int{}

	for key, element := range m {
		if element > temp {
			temp = element
			prevkey = key

			largerNumber = temp
			k = prevkey
		}

		if element == temp {
			largerNumber = element
			temp = element
			k = prevkey
		}
	}

	arr = append(arr, largerNumber)
	arr = append(arr, k)

	return arr
}

func reverseMap(m map[int]int) map[int]int {
	n := make(map[int]int, len(m))
	for k, v := range m {
		n[v] = k
	}
	return n
}

func main() {
	High("aa b")
}
