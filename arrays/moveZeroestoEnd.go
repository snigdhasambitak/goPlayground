package main

import "fmt"

func moveZeroes(nums []int) {
	maxZeros := 0
	l := len(nums)

	for i := range nums {

		if nums[i] != 0 {
			nums[maxZeros] = nums[i]
			maxZeros++
		}
	}

	for i := maxZeros; i < l; i++ {
		nums[i] = 0
	}
	fmt.Println(nums)

}
func main() {

	nums := []int{0, 0, 1}
	moveZeroes(nums)
}
