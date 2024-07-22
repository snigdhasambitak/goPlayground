package main

import "fmt"

func main() {

	// pre allocate the space using the make function
	names := make([]string, 3)

	// assign the values
	names[0] = "ken"
	names[1] = "kenny"
	names[2] = "kennith"

	fmt.Println(names)

}
