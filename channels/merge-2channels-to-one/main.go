package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func mergeChannels(ctx context.Context, chs ...chan int) <-chan int {
	mergeCh := make(chan (int))
	wg := sync.WaitGroup{}

	for _, ch := range chs {
		wg.Add(1)

		go func(ch chan int) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case el, ok := <-ch:
					if ok {
						mergeCh <- el
					}
				}
			}

		}(ch)
	}

	// for _, ch := range chs {
	// 	wg.Add(1)

	// 	go func(ch chan int) {
	// 		defer wg.Done()

	// 		for el := range ch {
	// 			mergeCh <- el
	// 		}
	// 	}(ch)
	// }

	go func() {
		defer close(mergeCh)
		wg.Wait()
	}()

	return mergeCh
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	ch1 := make(chan (int))
	ch2 := make(chan (int))
	ch3 := make(chan (int))

	go func() {
		for _, el := range []int{1, 2, 3} {
			ch1 <- el
		}

		close(ch1)
	}()

	go func() {
		for _, el := range []int{4, 5, 6} {
			ch2 <- el
		}

		close(ch2)
	}()

	go func() {
		for _, el := range []int{40, 50, 60} {
			ch3 <- el
		}

		close(ch3)
	}()

	for el := range mergeChannels(ctx, ch1, ch2, ch3) {
		fmt.Printf("%d %s", el, " , ")
	}
}
