package main

import (
	"fmt"
	"math"
)

func main() {
	a := 9.1
	fmt.Printf("the sq root of %d is %g \n", a, math.Sqrt(a))
	fmt.Println(math.Ceil(a))
	fmt.Println(math.Pow(9.0, 3))
}
