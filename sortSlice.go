package main

import (
	"fmt"
	"sort"
)

func main() {

	arr := []int{89, 1, 65, 2, 7, 0, 45, 54}

	// sort in ascending order
	sort.Ints(arr)
	fmt.Println(arr)

	//sort descending order
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})
	fmt.Println(arr)

}
