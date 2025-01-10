package main

import "fmt"

// permute generates all permutations of the given array.
func permutation(arr []int) [][]int {
	var result [][]int
	var generatePermutations func([]int, int)

	generatePermutations = func(arr []int, start int) {
		if start == len(arr) {
			temp := make([]int, len(arr))
			copy(temp, arr)
			result = append(result, temp)
			fmt.Println(result)
			return
		}

		// backtracking
		for i := start; i < len(arr); i++ {
			//fmt.Printf("i is %d, start is %d, arr is %d\n", i, start, arr)
			arr[start], arr[i] = arr[i], arr[start]
			generatePermutations(arr, start+1)
			arr[start], arr[i] = arr[i], arr[start]
			//fmt.Printf("i after function is %d, start after function is %d, arr is %d\n", i, start, arr)
		}
	}

	generatePermutations(arr, 0)
	return result
}

func main() {
	arr := []int{1, 2, 3, 4}
	permutations := permutation(arr)

	fmt.Println("All permutations:")
	for _, perm := range permutations {
		fmt.Println(perm)
	}
}
