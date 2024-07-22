package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s is %d years old", p.Name, p.Age)
}

func main() {
	a := Person{"Alice", 20}
	b := Person{"Bob", 30}
	fmt.Println(a, b)
}
