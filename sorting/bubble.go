package main

import "fmt"

func bubbleSort(arr []int) []int {
	n := len(arr)
	if n <= 1 {
		return arr
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}

func main() {
	arr := []int{8, 9, 7, 1, 3, 6}
	fmt.Println(bubbleSort(arr))
}
