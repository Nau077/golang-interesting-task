package main

import (
	"fmt"
)

type sema chan struct{}

func Add(n int) sema {
	return make(sema, n)
}

func (s sema) Done() {
	s <- struct{}{}
}

func (s sema) Wait(k int) {
	for i := 0; i < k; i++ {
		<-s
	}
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	n := len(numbers)

	sem := Add(n)
	defer close(sem)

	for _, num := range numbers {
		go func(n int) {
			defer sem.Done()

			fmt.Println(n)
		}(num)
	}

	sem.Wait(n)
	fmt.Println("after wait")

}
