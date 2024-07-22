package main

import (
	"fmt"
	"runtime"
	"time"
)

func osSwitch(os string) {

	fmt.Print("Go runs on ")
	switch os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s", os)
	}
}

func daySwitch(day time.Weekday) {
	switch time.Saturday {
	case day + 0:
		fmt.Println("Today is Saturday")
	case day + 1:
		fmt.Println("Tomorrow is Saturday")
	case day + 2:
		fmt.Println("Day after tomorrow is Saturday")
	default:
		fmt.Println("Saturday is far")
	}

}

func noConditionSwitch() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Morning")
	case t.Hour() < 17:
		fmt.Println("Afternooon")
	default:
		fmt.Println("Evening")
	}

}

func age() {

	ages := map[string]int{}
	ages["kevin"] = 7

	switch ages["kevin"] {

	case 1, 3, 5, 7, 11:
		fmt.Println("kevin age is prime")
	case 2, 4, 6, 8:
		fmt.Println("kevin age is even")
	default:
		fmt.Println("kevin age does not matter")

	}

}

func main() {
	// determine osx
	os := runtime.GOOS
	osSwitch(os)

	// determine day
	day := time.Now().Weekday()
	daySwitch(day)

	// no condition switch
	noConditionSwitch()

	// switch condition
	age()

}
