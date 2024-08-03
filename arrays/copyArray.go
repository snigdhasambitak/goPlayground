package main

import "fmt"

func main() {
	seq := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

	// copy array
	var seq1 []int

	seq1 = append(seq1, seq...)

	// remove index 6
	seq1 = append(seq1[:6], seq1[6+1:]...)
	fmt.Println(seq1)

}
