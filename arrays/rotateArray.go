package main

import "fmt"

func rotateArray(nums []int, k int) {

	for i := 0; i < k; i++ {

		v := nums[len(nums)-1]
		nums = append([]int{v}, nums[:len(nums)-1]...)
	}
	fmt.Println(nums)

}

func rotateArray2(nums []int, k int) {

	lengthofArray := len(nums)

	reverse(nums, 0, lengthofArray-1)
	reverse(nums[:k], 0, k-1)
	reverse(nums[k:], 0, lengthofArray-k-1)

	fmt.Println(nums)

}

func reverse(nums []int, start, end int) {

	for i, j := start, end; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func main() {

	nums := []int{1, 2, 3, 4, 5, 6, 7}

	k := 3

	rotateArray(nums, k)
	rotateArray2(nums, k)

}
