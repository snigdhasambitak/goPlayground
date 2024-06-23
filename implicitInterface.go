package main

import (
	"fmt"
	"math"
)

type Inter interface {
	methodInter()
}

type f float64

type Str struct {
	name string
}

// This method means type T implements the interface I,
func (s *Str) methodInter() {
	fmt.Println(s)
}

func (f f) methodInter() {
	fmt.Println(f)
}

func main() {
	var I Inter
	I = &Str{"hello"}
	describe(I)
	I.methodInter()

	I = f(math.Pi)
	describe(I)
	I.methodInter()

}

func describe(i Inter) {
	fmt.Printf("(%v, %T)\n)", i, i)
}
