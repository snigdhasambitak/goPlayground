package main

import "fmt"

func channel2(a []int, c chan int) {

	for _, n := range a {
		c <- n
	}
	close(c)
}

func main() {
	c := make(chan int)

	var arr2 = []int{9, 0, 3, 1, 8, 2, 10, 78, 54, 91}

	go channel2(arr2, c)
	for i := range c {
		fmt.Println(i)
	}
}
