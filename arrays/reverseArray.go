package main

import (
	"fmt"
	"sort"
)

func main() {

	var arr1 = []int{10, 2, 3, 9, 1, 7, 32, 91, 42}

	sort.Sort(sort.Reverse(sort.IntSlice(arr1)))
	fmt.Println(arr1)

	sort.Ints(arr1)
	fmt.Println(arr1)

	sort.Slice(arr1, func(i, j int) bool {
		return arr1[i] > arr1[j]
	})
	fmt.Println(arr1)
}
