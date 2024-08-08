package main

import "fmt"

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	fmt.Println(n)
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	fmt.Println(fibonacci(8))

}
