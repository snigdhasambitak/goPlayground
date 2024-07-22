package main

import "fmt"

type I3 interface {
}

type Mtype struct {
	name string
}

func main() {
	var i I3
	i = Mtype{"little"}
	describe3(i)

	var i4 interface{}
	i4 = "title"
	describe3(i4)

	i4 = 10
	describe3(i4)

}

func describe3(i I3) {
	fmt.Printf("(%v, %T)\n", i, i)
}
