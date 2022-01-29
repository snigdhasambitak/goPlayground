package main

import(
	"fmt"
	"math"
)

func sqrt(a float64) float64 {
	if a < 0 {
		return math.Sqrt(-a)
	}
	return math.Sqrt(a)
}

func pow(n, m, lim float64) float64 {
	if v:= math.Pow(n, m); v < lim{
		return v
	} else {
		fmt.Printf("%g >= %g\n",v, lim)
	}
	return lim
}

func main(){
	fmt.Println(sqrt(-10))
	fmt.Println(
		pow(3,2,10),
		pow(3,3,20),
		)

}