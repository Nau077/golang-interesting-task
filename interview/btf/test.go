package main

import (
	"fmt"
	"sync"
)

func main() {
	test()
}

func test() {
	var wg sync.WaitGroup
	wg.Add(3)
	c := make(chan int, 3)
	b := make(chan int)

	go func() {
		defer wg.Done()
		c <- 1

		close(c)
	}()

	go func() {
		defer wg.Done()
		b <- 2

		close(b)
	}()

	go func() {
		defer wg.Done()

		for el := range b {
			fmt.Println(el)
		}
	}()

	wg.Wait()
}
