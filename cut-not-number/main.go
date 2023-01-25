package main

import (
	"fmt"
	"regexp"
)

var nonAlphanumericRegex = regexp.MustCompile(`[^0-9 ]+`)

func main() {
	str := "mfgah134517095aldrfgvh8h"
	b := nonAlphanumericRegex.ReplaceAllString(str, "")

	fmt.Println(b)
}
