package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	// Create a reader from a string
	r := strings.NewReader("This is a sunny day!")

	// Create a buffer to hold the read data
	b := make([]byte, 4) // buffer size of 4 bytes
	for {
		n, err := r.Read(b) // The Read method is called in a loop to read data into the buffer

		if err != nil {
			if err == io.EOF {
				fmt.Println("EOF reached")
				break
			}
			fmt.Println("Error:", err)
			break

		}

		fmt.Printf("Read %d bytes: %s\n", n, b[:n])

	}
}
