package main

import (
	"fmt"
	"time"
)

func main() {
	results := make(chan int)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	for _, num := range numbers {
		go func(num int) {
			result := num * num
			time.Sleep(1 * time.Second)
			results <- result
		}(num)
	}

	for range numbers {
		fmt.Println(<-results)
	}

}
