package main

import "fmt"

var i, j int = 2, 3

func swap(a, b string) (string, string) {
	return b, a
}


func swapnum(a, b int) (x, y int){
	x, y = b, a
	return
}

func main(){
	a, b := "hello", "world"
	fmt.Println(swap(a, b))
	fmt.Println(swapnum(5,9))
	var c, d = false, true
	fmt.Println(i, j, c, d)
}