package main

import "fmt"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Length, Breadth float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Breadth
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius
}

func printArea(shape Shape) {
	fmt.Println(shape.Area())
}

func main() {
	r := Rectangle{10, 20}
	c1 := Circle{10}
	printArea(r)
	printArea(c1)

}
