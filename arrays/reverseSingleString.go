package main

import (
	"fmt"
)

func main() {
	s := "Hi My name is Aryakumar"

	// Convert the string to a slice of runes to handle Unicode characters
	s1 := []rune(s)

	fmt.Println(string(s1))

	// Print each rune (character)
	for i := 0; i < len(s1); i++ {
		fmt.Printf("%c\n", s1[i])
	}

	// reverse the rune slice
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s1[i], s1[j] = s1[j], s1[i]
	}

	fmt.Println(string(s1))

}
