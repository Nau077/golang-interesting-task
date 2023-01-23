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
		for i := 1; i < 10; i++ {
			inCh <- i
		}
		wg.Done()
		close(inCh)
	}()

	go func() {
		for v := range inCh {
			outCh <- v * v
		}
		wg.Done()
		close(outCh)
	}()

	go func() {
		for v := range outCh {
			fmt.Println(v)
		}
		wg.Done()
	}()
	wg.Wait()
}
