package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("hello world\n")
	fmt.Println(time.Now(), "\n")
	Hours := 7
	Mins := 30
	Sec := 30
	timein := time.Now().Local().Add(time.Hour*time.Duration(Hours) +
		time.Minute*time.Duration(Mins) +
		time.Second*time.Duration(Sec))
	fmt.Println(timein)

}
