package main

import (
	"fmt"
	"strings"
)

func main() {

	s := "this is a palindrome strong"

	fmt.Println(s)
	b := strings.Fields(s)

	fmt.Println(b)

	sm := make(map[string]int)

	for _, v := range b {

		sm[v]++
	}
	fmt.Println(sm)
}
