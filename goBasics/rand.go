package main

import (
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	min := 20
	max := 40
	println(rand.Intn( max - min + 1 ) + min)
}
