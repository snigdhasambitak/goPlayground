package main

import (
	"fmt"
	"os"
)

func main() {

	// use defer with count
	defercount()

	// defer file example
	deferfile()

}

func defercount() {
	fmt.Println("count start")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}

func deferfile() {
	f, err := os.Create("file")
	if err != nil {
		panic("cannot create file")
	}
	defer f.Close()
	fmt.Println("file closed")

}
