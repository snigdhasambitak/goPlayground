package main

import "fmt"

type Vertex3 struct {
	X, Y float64
}

func (v Vertex3) Abs() float64 {
	return v.X*v.X + v.Y*v.Y
}

func AbsFunc(v Vertex3) float64 {
	return v.X*v.X + v.Y*v.Y
}

func main() {
	v := Vertex3{3, 4}

	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	p := &Vertex3{4, 3}
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p))
}
