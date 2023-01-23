package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func randomNumberGenerate(ctx context.Context, inCh chan int, outCh chan int, wg *sync.WaitGroup) {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for {
		select {
		case <-ctx.Done():
			close(inCh)
			wg.Done()
			return
		case n, ok := <-inCh:
			if ok {
				outCh <- r.Intn(n)
			}
		}
	}
}

func main() {
	inCh := make(chan int)
	outCh := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())

	var wg sync.WaitGroup
	wg.Add(2)

	go randomNumberGenerate(ctx, inCh, outCh, &wg)
	go func() {

		// for {
		// 	select {
		// 	case <-ctx.Done():
		// 		wg.Done()
		// 		return
		// 	case n, ok := <-outCh:
		// 		if ok {
		// 			fmt.Println(n)
		// 		}
		// 	}
		// }

		for n := range outCh {
			time.Sleep(time.Second)
			fmt.Println(n)
		}

		wg.Done()
		close(outCh)
	}()

	for i := 5; i < 7; i++ {
		inCh <- i
	}

	cancel()
	wg.Wait()

}
