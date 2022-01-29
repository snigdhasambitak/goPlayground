package main

import (
	"fmt"
	"os"
)

func main(){

	f, err := os.Create("file")
	if err != nil{
		panic("cannot create file")
	}
	defer f.Close()
	fmt.Println("file closed")
	
}
