package main

import "fmt"

func rotateArray(nums []int, k int) {

	for i := 0; i < k; i++ {

		v := nums[len(nums)-1]
		nums = append([]int{v}, nums[:len(nums)-1]...)
	}
	fmt.Println(nums)

}

func main() {

	nums := []int{1, 2, 3, 4, 5, 6, 7}

	k := 3

	rotateArray(nums, k)

}
