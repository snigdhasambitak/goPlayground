package main

import "math"

func sqrt(a float64) float64 {
	if a < 0{
		return (math.Sqrt(-a))
	}
	return (math.Sqrt(a))
}

func main(){
	println(sqrt(-20))
}