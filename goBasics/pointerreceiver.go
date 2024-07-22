package main

import (
	"fmt"
	"math"
)

type Vertex4 struct {
	X, Y float64
}

func (v *Vertex4) Scale(f float64) {
	v.X = f * v.X
	v.Y = f * v.Y
}

func (v *Vertex4) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex4{3, 4}

	fmt.Printf("Before scale: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(2)
	fmt.Printf("After scale: %+v, Abs: %v\n", v, v.Abs())
}
