package main

import (
	"fmt"
	"sort"
)

func matchPairs(array []int, target int) [][]int {

	i := 0
	j := len(array) - 1

	result := make([][]int, 0)
	sort.Ints(array)
	for i < j {

		if array[i]+array[j] == target {
			result = append(result, []int{array[i], array[j]})
			i++
			j--
		} else if array[i]+array[j] > target {
			j--
		} else {
			i++
		}
	}
	return result

}

func main() {
	array := []int{10, 1, 4, 2, 3, 9, 5, 6, 7, 8}
	fmt.Printf("%v\n", matchPairs(array, 10))

}
