package main

import "fmt"

func main(){
	sum := 0
	for i:=0; i < 10; i++ {
		if i% 2 == 0{
			fmt.Printf("%v is a Even number \n", i)
		}
		sum += i
	}
	fmt.Println(sum)

	for sum < 91 {
		sum += sum
	}
	fmt.Println(sum)
}
