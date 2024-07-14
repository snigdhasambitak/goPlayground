package main

import "fmt"

func SlicetoMap(arr []int) map[int]struct{} {
	m := make(map[int]struct{})
	for _, v := range arr {
		m[v] = struct{}{}
	}
	return m
}

func FindElement(m map[int]struct{}, ele int) bool {
	_, ok := m[ele]
	return ok
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	m := SlicetoMap(arr)
	fmt.Println(m)
	fmt.Println(FindElement(m, 3))
	fmt.Println(FindElement(m, 8))
}
