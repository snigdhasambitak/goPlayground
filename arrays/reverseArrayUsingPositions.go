package main

import "fmt"

func reverseArrayUsingPositions(nums []int) []int {
	lenArr := len(nums) - 1
	for i := 0; i < lenArr; {

		nums[i], nums[lenArr] = nums[lenArr], nums[i]
		i++
		lenArr--
	}
	return nums
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(reverseArrayUsingPositions(nums))
}
