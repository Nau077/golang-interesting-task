package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomNumberGenerate(chLimit int) <-chan int {
	ch := make(chan int)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	go func() {
		for i := 5; i < chLimit; i++ {
			ch <- r.Intn(i)
		}
		close(ch)
	}()

	return ch
}

func main() {

	ch := randomNumberGenerate(10)

	for n := range ch {
		time.Sleep(time.Second)
		fmt.Println(n)
	}
}
