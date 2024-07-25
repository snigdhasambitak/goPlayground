package main

import (
	"fmt"
	"sort"
)

func maxSubArrayWorstCase(nums []int) int {

	arrsum := make([]int, 0)

	for i := range nums {
		sum := nums[i]
		for j := i + 1; j < len(nums); j++ {
			sum += nums[j]
			arrsum = append(arrsum, sum)
		}
	}

	fmt.Println(arrsum)

	sort.Ints(arrsum)

	return arrsum[len(arrsum)-1]
}

func maxSubArrayBestCase(nums []int) int {
	currentSum := nums[0]
	Largest := nums[0]

	for i := 1; i < len(nums); i++ {

		if nums[i] > currentSum {
			currentSum = nums[i]
		} else {
			currentSum += nums[i]
		}
		if currentSum > Largest {
			Largest = currentSum
		}
	}
	return Largest
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArrayWorstCase(nums))
	fmt.Println(maxSubArrayBestCase(nums))
}
