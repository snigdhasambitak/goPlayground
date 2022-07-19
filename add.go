package main

import (
	"fmt"
)

func add(a float64, b float64) float64 {
	return a + b
}

func sub(a int, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}

}

func mul(a, b int) int {
	return a * b
}

func divide(a, b float64) float64 {
	return a / b
}

func main() {
	fmt.Println(add(1.5, 8))
	fmt.Println(sub(2, 8))
	fmt.Println(mul(3, 5))
	fmt.Println(divide(7, 2))
}
