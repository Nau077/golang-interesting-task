package main

import (
	"fmt"
	"sync"
)

func mergeChannels(ch ...<-chan int) <-chan int {
	mergeCh := make(chan int)

	wg := sync.WaitGroup{}
	wg.Add(len(ch))

	for i := 0; i < len(ch); i++ {
		go func(i int) {
			defer wg.Done()

			ch := ch[i]

			for {
				select {
				case el, ok := <-ch:
					if ok {
						mergeCh <- el
					} else {
						return
					}
				}
			}

			// for el := range ch {
			// 	mergeCh <- el
			// }
		}(i)
	}

	go func() {
		wg.Wait()
		close(mergeCh)
	}()

	return mergeCh
}

func main() {
	a := make(chan int)
	b := make(chan int)
	c := make(chan int)

	go func() {
		defer close(a)
		for i := 0; i < 10; i++ {
			a <- i
		}
	}()

	go func() {
		defer close(b)
		for i := 10; i < 20; i++ {
			b <- i
		}
	}()

	go func() {
		defer close(c)
		for i := 20; i < 40; i++ {
			c <- i
		}
	}()

	for el := range mergeChannels(a, b, c) {
		fmt.Printf("%d, ", el)
	}
}
