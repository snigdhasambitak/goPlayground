package main

import "fmt"

func matchingPairs(array []int, targetSum int) [][]int {

	j := make(map[int]bool)

	result := make([][]int, 0)
	for _, v := range array {
		potential := targetSum - v

		_, ok := j[potential]
		if ok {
			result = append(result, []int{v, potential})
		}
		j[v] = true
	}
	return result

}

func main() {
	array := []int{10, 1, 4, 2, 3, 9, 5, 6, 7, 8}
	fmt.Printf("%v\n", matchingPairs(array, 10))

}
