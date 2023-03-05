package main

func main() {
	checkFirst()
}

func checkFirst() {
	chWrite := make(chan int)
	chWriteD := make(chan int, 2)

	//go func() {
	chWrite <- 556
	//}()

	go func() {
		chWriteD <- 556
	}()
}
