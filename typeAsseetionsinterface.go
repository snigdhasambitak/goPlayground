package main

import "fmt"

func main() {

	var I interface{} = "hello"

	i := I.(string)
	fmt.Println(i)

	s, ok := I.(string)
	fmt.Println(s, ok)

	j, ok := I.(int)
	fmt.Println(j, ok)

	k := I.(float64)
	fmt.Println(k)

}
