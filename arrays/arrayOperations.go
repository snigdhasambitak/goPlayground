package main

import "fmt"

func arrayCopy(arr []int) []int {

	// arrCopy := make([]int, len(arr))
	//copy(arrCopy, arr)

	arrCopy := make([]int, 0)
	arrCopy = append(arr, arrCopy...)

	return arrCopy
}

func arrayPush(arr []int, value int) []int {
	arr = append(arr, value)
	return arr
}

func arrayPushBeginningShift(arr []int, value int) []int {
	return append([]int{value}, arr[:len(arr)-1]...)
}

func arrayPop(arr []int) []int {
	return arr[:len(arr)-1]
}

func arrayPopIndex(arr []int, index int) []int {
	return append(arr[:index], arr[index+1:]...)
}

func arrayPopBeginning(arr []int) []int {
	return arr[1:]
}

func arrayInsert(arr []int, value int, index int) []int {
	return append(arr[:index], append([]int{value}, arr[index:len(arr)-1]...)...)
}

func main() {

	// arrayCopy
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(arrayCopy(arr))

	//arrayPush [1 2 3 4 5 6]
	fmt.Println(arrayPush(arr, 6))

	//arrayPushBeginning [9 1 2 3 4 5]
	fmt.Println(arrayPushBeginningShift(arr, 9))

	//arrayInsert [1 2 12 3 4]
	fmt.Println(arrayInsert(arr, 12, 2))

	//arrayPop [1 2 3 4]
	fmt.Println(arrayPop(arr))

	//arrayPopBeginning [2 3 4 5]
	fmt.Println(arrayPopBeginning(arr))

	//arrayPopIndex [1 2 3 5]
	fmt.Println(arrayPopIndex(arr, 3))

}
