package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	b := [4]int{1, 2, 3, 4}

	var c []int

	c = append(c, b[0:4]...)

	sumofarray(c)
}

func sumofarray(a ...[]int) {
	for i, j := range a {
		fmt.Println(i, j)
	}
}
