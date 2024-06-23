package main

import "fmt"

type Inter2 interface {
	method2()
}

type T2 struct {
	S2 string
}

func (t *T2) method2() {
	if t == nil {
		fmt.Println("t is nil")
		return
	}
	fmt.Println(t.S2)

}

func main() {
	var i Inter2
	var t *T2

	i = t
	describe2(i)
	i.method2()

	i = &T2{"hello"}
	describe2(i)
	i.method2()
}

func describe2(i Inter2) {
	fmt.Printf("(%v, %T)\n", i, i)
}
