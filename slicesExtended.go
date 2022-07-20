package main

import "fmt"

func main() {

	s := []int{1, 2, 3, 4, 5}
	printslice(s)

	s = s[:2]
	printslice(s)

	s = s[:4]
	printslice(s)

	b := s[2:4]
	printslice(b)

	s = s[1:5]
	printslice(s)

	var a []int

	nilslice(a)
}

func printslice(s []int) {

	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func nilslice(a []int) {
	if a == nil {
		fmt.Println("nil")
	}
}
