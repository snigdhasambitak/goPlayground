package main

import "fmt"

func main(){
	i, j := 21, 301

	// points to i and j
	p := &i
	q := &j

	// print the values of i and j via pointer
	fmt.Println(*p)
	fmt.Println(*q)

	// change the value of i and via pointer
	*p = 42
	*q = 201
	fmt.Println(i)
	fmt.Println(j)

	// change the value of j via pointer
	p = &j
	*p = *p * 2
	fmt.Println(j)



}
