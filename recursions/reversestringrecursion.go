package main

import "fmt"

func reverseString(s string) string {
	if len(s) < 2 {
		return s
	}

	return reverseString(s[1:]) + string(s[0])
}

func main() {
	fmt.Println(reverseString("hello"))
}
