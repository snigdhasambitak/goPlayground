package main

import "fmt"

func BinarySearch(array []int, target int) int {
	i, j := 0, len(array)

	for i < j {
		mid := i + (j-i)/2
		fmt.Println(i, j, mid)
		if array[mid] == target {
			return mid
		}
		if array[mid] < target {
			i = mid + 1
		} else {
			j = mid - 1
		}

	}
	return -1
}

func main() {
	array := []int{0, 1, 21, 33, 45, 45, 61, 71, 72, 73}
	fmt.Println(BinarySearch(array, 33))
}
