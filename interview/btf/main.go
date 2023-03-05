package main

import (
	"fmt"
)

func main() {
	fmt.Println(test())
}

func test() int {
	var x int
	x = 9
	defer func(n int) {
		fmt.Println(n)
		fmt.Println(x)

		x *= 10
	}(x)

	return x

}
