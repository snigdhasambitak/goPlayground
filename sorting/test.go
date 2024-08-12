package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{10, 7, 9, 1, 5}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	fmt.Println(arr)

	sort.Sort(sort.Reverse(sort.IntSlice(arr)))
	fmt.Println(arr)

	chars := []rune{'e', 'g', 'o', 'a', 'c'}

	sort.Slice(chars, func(i, j int) bool {
		return chars[i] < chars[j]
	})
	strings := make([]string, len(chars))
	for i, char := range chars {
		strings[i] = string(char)
	}
	fmt.Println(strings)
}
