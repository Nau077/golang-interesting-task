package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	inCh := make(chan int)
	outCh := make(chan int)
	wg.Add(3)

	go func() {
		defer wg.Done()
		for i := 1; i < 10; i++ {
			inCh <- i
		}

		close(inCh)
	}()

	go func() {
		defer wg.Done()
		for v := range inCh {
			outCh <- v * v
		}

		close(outCh)
	}()

	go func() {
		defer wg.Done()
		for v := range outCh {
			fmt.Println(v)
		}
	}()
	wg.Wait()
}
