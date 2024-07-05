package main

import "fmt"

type Numeric interface {
	~int | ~float64
}

func add[T Numeric](a, b T) T {
	return a + b
}

func Swap[T any](x, y T) (T, T) {
	return y, x
}

func main() {

	// swap function using type constraints
	x, y := 1, 2
	fmt.Println(Swap(x, y))

	s1, s2 := "hi", "hello"
	fmt.Println(Swap(s1, s2))

	// add using type constraint interface
	a, b := 3.1, 4.3
	fmt.Println(add(a, b))

}
