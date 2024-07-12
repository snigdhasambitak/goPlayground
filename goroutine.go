package main

import (
	"fmt"
	"time"
)

func say(s string) {
	fmt.Println(s)
}

func printNumbers() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		time.Sleep(500 * time.Millisecond)
	}
}

func printChars() {
	for i := 'a'; i < 'e'; i++ {
		fmt.Printf("%c\n", i)
		time.Sleep(500 * time.Millisecond)
	}
}
func main() {
	go say("hello")
	time.Sleep(1 * time.Second)

	go printNumbers()
	go printChars()
	time.Sleep(5 * time.Second)

}
