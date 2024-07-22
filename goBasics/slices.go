package main

import "fmt"

func main() {

	// define an array
	primes := [5]int{1, 3, 5, 7, 11}

	// define a slice
	var a []int

	// assign elements of an array to a slice
	a = primes[1:4]

	// use append function to append elements
	a = append(a, 10) // duplicates the list of elements in a
	fmt.Println(a)

	// reassigning an element in a slice will change the actual value in array
	a[0] = 1

	// print the actual primes
	fmt.Println(primes)

	// define an array of strings
	names := [4]string{"amit", "aman", "amar", "amul"}

	nameSlice1 := names[1:3]
	nameSlice2 := names[0:3]

	fmt.Println(nameSlice1, nameSlice2)

	nameSlice1[0] = "xxx"

	fmt.Println(nameSlice1, nameSlice2)
	fmt.Println(names)

}
