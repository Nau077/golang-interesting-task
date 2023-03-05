package main

import (
	"fmt"
)

func main() {
	fmt.Println(test())
}

func test() int {
	var x int
	x = 3
	defer func(n int) {
		fmt.Println(n)
		fmt.Println(x)

		x *= 10
	}(x)

	x = 8

	return 6
}

// 3
// 8
// 6
