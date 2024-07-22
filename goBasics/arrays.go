package main

import "fmt"

func main() {
	var a [2]int
	a[0] = 1
	a[1] = 2
	fmt.Println(a)

	var b [2]string
	b[0] = "hello"
	b[1] = "world"
	fmt.Println(b[1], b[0])

	prime := [5]int{1, 3, 5, 7, 11}
	fmt.Println(prime)

	name := [4]string{"amar", "aman", "amit"} // dynamic assignment
	fmt.Println(name)

	fmt.Println("name[3] is nil :", name[3] == "") // default type of string is " "

	booltype := [2]bool{true, false}

	fmt.Println(booltype)

}
