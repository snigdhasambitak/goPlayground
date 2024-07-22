package main

import "fmt"

type Vertex struct {
	X, Y int
}

type Pod struct {
	Image, Version string
}

func main(){
	v := Vertex{1,2}
	v.X = 10
	v.Y = 20
	fmt.Println(v)

	// pointer to struct
	p := &v
	p.X = 100
	p.Y = 200
	fmt.Println(v)

	// Implicit values
	v1 := Vertex{X:1}
	v2 := Vertex{Y:2}
	p1 := & Vertex{5,6}
	fmt.Println(v1, v2, p1)

	// string struct
	s1 := Pod{"nginx", "1.0"}
	s2 := &s1
	s2.Image = "busybox"
	s2.Version = "2.0"
	fmt.Println(s1, s2)
}
