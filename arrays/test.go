package main

import (
	"fmt"
	"strings"
)

func main() {

	s := "This is a laptop 123"

	sarr := []rune(s)

	for i, j := 0, len(sarr)-1; i < j; i, j = i+1, j-1 {
		sarr[i], sarr[j] = sarr[j], sarr[i]
	}
	fmt.Println(string(sarr))

	m := "My name is lakhan"

	marr := strings.Fields(m)
	fmt.Println(marr)

}
