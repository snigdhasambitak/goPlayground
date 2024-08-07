package main

import (
	"fmt"
	"sort"
)

func ThreeNumberSum(nums []int, target int) [][]int {
	marr := make([][]int, 0)
	sort.Ints(nums)
	for i := range nums {
		j, m := i+1, len(nums)-1
		for j < m {
			if nums[i]+nums[j]+nums[m] == target {
				marr = append(marr, []int{nums[i], nums[j], nums[m]})
				j++
				m--
			}
			if nums[i]+nums[j]+nums[m] < target {
				j++
			}
			if nums[i]+nums[j]+nums[m] > target {
				m--
			}
		}
	}

	if len(marr) > 0 {
		return marr
	} else {
		return [][]int{}
	}

}

func main() {
	arr, targetSum := []int{12, 3, 1, 2, -6, 5, -8, 6}, 0
	fmt.Println(ThreeNumberSum(arr, targetSum))

}
