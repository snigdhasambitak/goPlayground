package main

import "fmt"

func main() {

	n := 10
	fmt.Println(sumofn(n))

	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	fmt.Println(mapsum(m))

	stopiteration(10)

}

func sumofn(n int) int {
	sum := 0

	// loop to iterate over n numbers
	for i := 0; i < n; i++ {
		sum += i
	}
	return sum
}

func mapsum(m map[string]int) int {
	sum := 0

	// loop to iterate over the values of a map
	for _, value := range m {
		sum += value
	}
	return sum
}

// utilise continue and break
func stopiteration(n int) {
	a := 0
	for a < n {
		if a%2 == 0 {
			a++
			continue
		} else if a == 5 {
			break
		}
		fmt.Println(a)
		a++
	}
}
