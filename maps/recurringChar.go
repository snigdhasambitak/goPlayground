package main

import "fmt"

func recurringChar(nums []int) int {

	arrmap := make(map[int]bool)

	for _, num := range nums {
		_, ok := arrmap[num]
		if ok {
			return num
		} else {
			arrmap[num] = true
		}
	}
	return -1

}

func main() {
	arr := []int{2, 5, 5, 2}
	fmt.Println(recurringChar(arr))
}
