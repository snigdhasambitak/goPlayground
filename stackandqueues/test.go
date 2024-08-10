package main

import "fmt"

type Places struct {
	place  string
	weight []int
}

func main() {

	p := Places{}
	p.place = "rome"
	p.weight = []int{1, 2, 3}
	sum := 0

	for _, v := range p.weight {
		sum += v
	}
	fmt.Println(sum)

}
