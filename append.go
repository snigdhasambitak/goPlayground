package main

import "fmt"

func main() {

	// define an empty slice
	var emptysliceint []int

	// use append function for slice
	emptysliceint = append(emptysliceint, 1, 2, 3)

	fmt.Println(emptysliceint)

	emptyslicestring := []string{}

	emptyslicestring = append(emptyslicestring, "xxx", "yyy", "zzz")

	fmt.Println(emptyslicestring)

}
