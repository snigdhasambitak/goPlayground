package main

import "fmt"

func sameElementsinArray(arr1 []int, arr2 []int) bool {

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
	fmt.Println(sameElementsinArray(arr1, arr2))
}
