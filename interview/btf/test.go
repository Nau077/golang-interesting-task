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
	c := make(chan int, 1)
	b := make(chan int)

	go func() {
		defer wg.Done()
		c <- 7

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
	//	time.Sleep(time.Second) // делает вывод 2:7
	fmt.Println(<-c) // 7:2 но это undefined behiviour

	wg.Wait()
}
