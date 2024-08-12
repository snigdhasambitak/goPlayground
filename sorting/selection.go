package main

import "fmt"

func selectionSort(arr []int) []int {

	for i := 0; i < len(arr)-1; i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
	return arr
}

func main() {
	arr := []int{8, 9, 7, 1, 3, 6}
	fmt.Println(selectionSort(arr))
}
