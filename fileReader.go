package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	file, err := os.Open("example.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	fileSize := fileInfo.Size()
	fmt.Println(fileSize)

	buffer := make([]byte, fileSize)
	file.Read(buffer)
	fmt.Println(string(buffer))
	//buffer := make([]byte, fileSize)
	//_, err = file.Read(buffer)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(string(buffer))

}
