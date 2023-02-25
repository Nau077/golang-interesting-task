package main

import (
	"fmt"
)

func main() {

	bfCH := make(chan interface{}, 3)
	bfCH <- 1
	bfCH <- 2
	bfCH <- 3
	bfCH <- 4 // если закомментировать, то всё будет ок

	fmt.Println(<-bfCH)
	fmt.Println(<-bfCH)
	fmt.Println("hey") // fatal error: all goroutines are asleep - deadlock!

}
