package main

import (
	"fmt"
)

// permute generates all permutations of the given slice of strings.
func permute(arr []string) [][]string {
	var result [][]string
	var generatePermutations func([]string, int)

	generatePermutations = func(arr []string, start int) {
		if start == len(arr) {
			temp := make([]string, len(arr))
			copy(temp, arr)
			result = append(result, temp)
			return
		}

		for i := start; i < len(arr); i++ {
			arr[start], arr[i] = arr[i], arr[start]
			generatePermutations(arr, start+1)
			arr[start], arr[i] = arr[i], arr[start] // backtrack
		}
	}

	generatePermutations(arr, 0)
	return result
}

// Calculate total weight of a given permutation
func calculateWeight(order []string, weights map[string][]int) int {
	totalWeight := 0
	for i, place := range order {
		if i < len(weights[place]) {
			totalWeight += weights[place][i]
		}
	}
	return totalWeight
}

func main() {
	// Define places and their corresponding weights
	placesmap := map[string][]int{
		"Paris":  {10, 24, 5, 16, 76},
		"London": {7, 3, 41, 61, 9},
		"Russia": {11, 15, 3, 8, 1},
		"Egypt":  {32, 2, 50, 34, 21},
		"Korea":  {1, 9, 21, 30, 3},
	}

	// Fetch place names directly from the map keys
	places := make([]string, 0, len(placesmap))
	for place := range placesmap {
		places = append(places, place)
	}

	// Generate all permutations of the places
	permutations := permute(places)

	// Find the permutation with the maximum weight
	maxWeight := 0
	var bestOrder []string
	for _, perm := range permutations {
		currentWeight := calculateWeight(perm, placesmap)
		if currentWeight > maxWeight {
			maxWeight = currentWeight
			bestOrder = perm
		}
	}

	// Output the results
	fmt.Printf("Maximum weight is %d\n", maxWeight)
	fmt.Println("Order to visit places:", bestOrder)
}
