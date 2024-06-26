package main

import (
	"errors"
	"fmt"
)

func dividebyzero(a, b float64) (float64, error) {

	if b == 0 {
		return 0, errors.New("divide by zero")
	}
	return a / b, nil
}

func main() {

	results, err := dividebyzero(10, 5)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(results)
	}

	results, err = dividebyzero(10, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(results)
	}
}
