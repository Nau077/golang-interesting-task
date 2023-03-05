package main

import (
	"errors"
	"fmt"
)

func binarySearch(s []int, el int) (int, error) {
	low := 0
	high := len(s) - 1

	for low <= high {
		mid := (low + high) / 2
		rein := s[mid]

		switch {
		case rein == el:
			return mid, nil
		case rein > el:
			high = mid - 1
		case rein < el:
			low = mid + 1
		}
	}

	return 0, errors.New("no el")
}

func main() {
	s := []int{
		1, 2, 5, 7, 8, 451, 472, 521, 711,
	}

	el := 7

	fmt.Println(binarySearch(s, el))
}
