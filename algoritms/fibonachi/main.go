package main

import "fmt"

func fib(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	a, b := 0, 1

	for i := 2; i < n; i++ {
		c := a + b
		a = b
		b = c
	}

	return b
}

func main() {
	fmt.Println(fib(8))
}
