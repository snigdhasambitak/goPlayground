package main

import (
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
	if v := math.Pow(n, m); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func main() {

	// simple example of if else if

	age := map[string]int{}

	age["kevin"] = 67

	if age["kevin"] < 67 {
		fmt.Println("kevin has not yet reached retirement age yet as his age is :", age["kevin"])
	} else if age["kevin"] > 67 {
		fmt.Println("kevin has reached retirement age yet as his age is :", age["kevin"])
	} else {
		fmt.Println("not yet retired")
	}

	fmt.Println(sqrt(-10))
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

}
