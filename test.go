package main

import "fmt"

func main() {

	var arr1 = []int{1, 2, 3}

	var arr2 = []int{1, 2, 3}

	arr3 := make([]int, 0)

	sum := 0

	for i := 0; i < len(arr1); i++ {
		sum = arr1[i] + arr2[i]
		arr3 = append(arr3, sum)
	}
	fmt.Println(arr3)

}
