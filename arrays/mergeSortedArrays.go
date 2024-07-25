package main

import (
	"fmt"
	"sort"
)

func mergeSortedArrayWorstCase(nums1 []int, nums2 []int) []int {
	nums3 := make([]int, 0)

	nums3 = append(nums3, append(nums1, nums2...)...)

	sort.Ints(nums3)

	return nums3

}

func mergeSortedArrayBestCase(nums1 []int, nums2 []int) []int {
	nums3 := make([]int, 0, len(nums1)+len(nums2))
	i, j := 0, 0
	m, n := len(nums1), len(nums2)

	if m == 0 {
		return nums2
	}

	if n == 0 {
		return nums1
	}

	for i < m && j < n {
		if nums1[i] < nums2[j] {
			nums3 = append(nums3, nums1[i])
			i++
		} else {
			nums3 = append(nums3, nums2[j])
			j++
		}
	}

	if m > n {
		nums3 = append(nums3, nums1[n:]...)
	} else {
		nums3 = append(nums3, nums2[m:]...)
	}

	return nums3

}

func main() {

	nums1 := []int{0, 3, 4, 31}
	nums2 := []int{4, 6, 30}
	fmt.Println(mergeSortedArrayWorstCase(nums1, nums2))
	fmt.Println(mergeSortedArrayBestCase(nums1, nums2))

}
