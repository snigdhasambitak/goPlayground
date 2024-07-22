package main

import (
	"fmt"
	"math/cmplx"
)

var(
	a uint64 = 1<<64 - 1
	b bool = false
	c complex128 = cmplx.Sqrt(12i - 5)
)

const pi = 3.14

func main(){
	fmt.Printf("Type: %T, Value: %v \n", a, a)
	fmt.Printf("Type: %T, Value: %v \n", b, b)
	fmt.Printf("Type: %T, Value: %v \n", c, c)
	var(
		i int
		k float64
		j bool
		s string
	)
	const hello = "world"

	fmt.Printf("%v, %v, %v, %q, %v, %s \n", i, k, j, s, pi, hello)
}
