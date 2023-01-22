package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func randomNumberGenerate(ctx context.Context, inCh chan int, outCh chan int) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for {
		select {
		case <-ctx.Done():
			return
		case n, ok := <-inCh:
			if ok {
				outCh <- r.Intn(n)
			} else {
				close(inCh)
				return
			}
		}
	}
}

func main() {
	inCh := make(chan int)
	outCh := make(chan int)
	ctx := context.Background()

	go randomNumberGenerate(ctx, inCh, outCh)

	go func() {
		for n := range outCh {
			fmt.Println(n)
		}
	}()

	for i := 5; i < 15; i++ {
		inCh <- i
	}

}
