package main

import "fmt"

func main() {
	checkFirst()
}

func checkFirst() {
	chWrite := make(chan int)
	chWriteD := make(chan int, 2)

	go func() {
		fmt.Println(123)
		chWrite <- 556
	}()
	chWriteD <- 556
	// go func() {

	// }()
}
