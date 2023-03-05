package main

import (
	"errors"
	"fmt"
)

func binarySearch(s []int, el int) (int, error) {
	low := 0
	high := len(s) - 1

	for low <= high {
		midIndex := (low + high) / 2
		midEl := s[midIndex]

		switch {
		case midEl == el:
			return midIndex, nil
		case midEl > el:
			high = midIndex - 1
		case midEl < el:
			low = midIndex + 1
		}
	}

	return 0, errors.New("error")
}

func main() {
	s := []int{
		1, 2, 5, 7, 8, 451, 472, 521, 711,
	}

	el := 2

	fmt.Println(binarySearch(s, el))
}
