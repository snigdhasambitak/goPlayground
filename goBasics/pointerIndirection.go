package main

import "fmt"

type Vertex2 struct {
	X, Y float64
}

func (v *Vertex2) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex2, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex2{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)
	fmt.Println(v)

	p := &Vertex2{4, 3}
	p.Scale(4)
	ScaleFunc(p, 10)
	fmt.Println(p)
}
