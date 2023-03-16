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

		if el == midEl {
			return midIndex, nil
		}

		if midEl > el {
			high = midIndex - 1
		}

		if midEl < el {
			low = midIndex + 1
		}
	}

	return 0, errors.New("error")
}

func main() {
	s := []int{2, 65, 455, 484, 522, 3226, 12322, 53234}

	fmt.Println(binarySearch(s, 522))
}
