package main

import (
	"fmt"
	"strings"
)

func reverseStrings(s string) string {

	sarray := strings.Fields(s)

	reversedArray := make([]string, 0)

	for i, j := 0, len(sarray)-1; i < j; i, j = i+1, j-1 {
		sarray[i], sarray[j] = sarray[j], sarray[i]
	}

	for i := range sarray {

		rune := []rune(sarray[i])

		for i, j := 0, len(rune)-1; i < j; i, j = i+1, j-1 {
			rune[i], rune[j] = rune[j], rune[i]
		}
		reversedArray = append(reversedArray, string(rune))
	}
	return strings.Join(reversedArray, " ")
}

func main() {

	s := "Hi My name is Aryakumar"
	fmt.Println(reverseStrings(s))
}
