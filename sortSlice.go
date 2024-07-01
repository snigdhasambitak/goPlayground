package main

import (
	"fmt"
	"sort"
)

func main() {

	arr := []int{10, 7, 1, 48, 9, 22}

	fmt.Println(arr)

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})

	for _, v := range arr {
		fmt.Println(v)
	}

}
