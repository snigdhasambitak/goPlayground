package main

import (
	"fmt"
	"os"
)

func main(){

	// defer count example
	fmt.Println("count start")
	for i:=0; i<10; i++{
		defer fmt.Println(i)
	}

	fmt.Println("done")

	// defer file example
	f, err := os.Create("file")
	if err != nil{
		panic("cannot create file")
	}
	defer f.Close()
	fmt.Println("file closed")
}
