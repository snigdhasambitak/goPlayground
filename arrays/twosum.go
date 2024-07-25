package main

import "fmt"

func twoSum(nums []int, target int) []int {

	numsmap := make(map[int]int)

	for i := range nums {
		potential := target - nums[i]

		v, ok := numsmap[potential]
		if ok {
			return []int{v, i}
		} else {
			numsmap[nums[i]] = i
		}
	}
	return []int{}

}

func main() {

	nums := []int{2, 7, 11, 15}
	fmt.Println(twoSum(nums, 17)) // [0,3]

}
