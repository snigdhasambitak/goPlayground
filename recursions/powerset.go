package main

import (
	"fmt"
)

func getPowerSet(array []int) [][]int {
	powerSet := [][]int{{}} // Start with the empty set

	for _, elem := range array {
		for _, subset := range powerSet {
			var newSubset []int
			newSubset = append(newSubset, subset...) // Create a copy of the current subset
			newSubset = append(newSubset, elem)      // Add the current element to the new subset
			powerSet = append(powerSet, newSubset)   // Add the new subset to the power set
		}
	}
	return powerSet
}

func main() {
	array := []int{1, 2, 3}
	powerSet := getPowerSet(array)
	fmt.Println(powerSet)
}
