package main

import (
	"fmt"
	"strconv"
)

// Digital root is the recursive sum of all the digits in a number.
// Given n, take the sum of the digits of n.
// If that value has more than one digit,
// continue reducing in this way until a single-digit number is produced.
// The input will be a non-negative integer.

// digital_root(16)
// 1 + 6 = 7

// digital_root(942)
// 9 + 4 + 2 = 15 ... 1 + 5 = 6

func DigitalRoot(n int) int {

	str := strconv.Itoa(n)
	strLen := len([]rune(str))
	var arr = []int{}
	var result int

	if strLen == 1 {
		return n
	}

	for _, el := range str {
		number, _ := strconv.Atoi(string(el))
		arr = append(arr, number)
	}

	for _, el := range arr {
		result += el
	}

	return DigitalRoot(result)
}

// func DigitalRoot(n int) int {
// 	for n > 9 {
// 		n = (n / 10) + n%10
// 	}

// 	return n
// }

// func DigitalRoot(n int) int {
//     return (n - 1) % 9 + 1
// }

func main() {
	res := DigitalRoot(16)
	fmt.Println(res)
}
