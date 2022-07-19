package main

import "fmt"

func main() {

	fmt.Println("greater than", 1 > 2)              // false
	fmt.Println("greater than or equal to", 2 >= 2) //true
	fmt.Println("lesser than", 2 < 1)               //false
	fmt.Println("equivalent", 1.0 == 1)             //true
	fmt.Println("not equivalent", 2 != 2)           //false
}
