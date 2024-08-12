package main

import "fmt"

func binarySearch(arr []int, key int) int {
	start, end := 0, len(arr)
	for start < end {
		mid := (start + end) / 2
		if key == arr[mid] {
			return mid
		}
		if key < arr[mid] {
			end = mid
		} else {
			start = mid + 1
		}
	}
	return -1
}
func main() {
	fmt.Println(binarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 5))
}
