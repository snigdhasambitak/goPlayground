package main

import "fmt"

// check if the 2nd array contains a single element matching 1st array

func sameElementinArray(arr1 []int, arr2 []int) bool {

	a := make(map[int]bool)

	for i := range arr1 {
		a[arr1[i]] = true
	}
	fmt.Println(a)
	for j := range arr2 {
		if a[arr2[j]] {
			return true
		}
	}
	return false
}

func main() {
	arr1 := []int{1, 2, 3}
	arr2 := []int{4, 5, 3}
	fmt.Println(sameElementinArray(arr1, arr2))
}
