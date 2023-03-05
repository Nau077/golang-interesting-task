package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("somefile")
	mydata := []byte("all my data I want to write to a file")
	if err != nil {
		fmt.Println(err)
	}

	err = ioutil.WriteFile("somefile", mydata, 0777)
	if err != nil {
		// Если произошла ошибка выводим ее в консоль
		fmt.Println(err)
	}

	fmt.Print(string(data))
}
