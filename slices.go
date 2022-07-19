package main

import "fmt"

func main(){
	primes := [5]int{1,3,5,7,11}

	// define a slice
	var a []int

	// assign elements of an array to a slice
	a = primes[1:3]

	fmt.Println(a)

	names := [5]string{"john", "jerrry", "tom", "hon", "nick"}
	var b []string

	b = names[2:5]
	fmt.Println(b)

	// Changing the elements of a slice modifies the corresponding elements of its underlying array

	b[1] = "XXX"

	fmt.Println(b)
	fmt.Println(names)
}
