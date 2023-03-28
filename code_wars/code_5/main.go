package main

// package kata
import "strings"

type KeyU struct {
	count  int
	origin string
}

func FirstNonRepeating(str string) string {
	m := make(map[string]KeyU)
	// сделали мапу из элементов
	// где ключ - буква, значение структура с счетчиком и оригиналом регистра
	for _, el := range str {
		stEl := string(el)
		stLow := strings.ToLower(stEl)
		_, ok := m[stLow]

		if !ok {
			m[stLow] = KeyU{
				count:  1,
				origin: stEl,
			}

		} else {
			n := m[stLow].count + 1
			m[stLow] = KeyU{
				count:  n,
				origin: stEl,
			}
		}
	}

	var sortS = []string{}
	strLow := strings.ToLower(str)

	for _, letter := range strLow {
		for key, val := range m {
			if string(letter) == key {
				if val.count == 1 {
					sortS = append(sortS, val.origin)
				}
			}
		}
	}

	if len(sortS) >= 1 {
		return sortS[0]
	} else {
		return ""
	}
}

// smart solution
// package kata
// import (
// 	"strings"
// )
// func FirstNonRepeating(str string) string {
//     for _, c := range str {
//         if strings.Count(strings.ToLower(str), strings.ToLower(string(c))) < 2 {
// 	          return string(c)
// 	      }
//     }
//     return ""
// }
