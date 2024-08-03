package main

import (
	"fmt"
)

func reverseArray(nums []int) []int {

	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]

	}
	return nums

}

func main() {
	arr := []int{7, 1, 2, 3, 4, 5, 6}
	fmt.Println(reverseArray(arr))

}
