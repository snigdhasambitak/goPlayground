package main

import "fmt"

func findDuplicates(nums []int) bool {

	arrmap := make(map[int]bool)

	for i := range nums {
		arrmap[nums[i]] = true
	}

	if len(arrmap) != len(nums) {
		return true
	}

	return false

}

func main() {
	nums := []int{1, 2, 3, 4, 5, 1, 3}
	fmt.Println(findDuplicates(nums))

}
