package main

import(
	"fmt"
	"runtime"
	"time"
)

func osSwitch(os string){

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
	fmt.Println(day)
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

func main(){
	// determine osx
	os:=runtime.GOOS
	osSwitch(os)

	// determine day
	day:=time.Now().Weekday()
	daySwitch(day)

}