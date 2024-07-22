package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v", err)
	}
	defer file.Close()

	reader, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("Error reading file: %v", err)
	}

	fmt.Println(string(reader))
}
