package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
)

func main() {

	file, err := os.Open("logfile.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r, err := regexp.Compile(`\b\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}\b`)
	if err != nil {
		log.Fatal(err)
	}

	fileline := bufio.NewScanner(file)

	logmap := make(map[string]int)

	for fileline.Scan() {
		text := fileline.Text()
		matches := r.FindAllString(text, -1)

		for _, m := range matches {
			logmap[m]++
		}

	}
	fmt.Println(logmap)

	type arrstruct struct {
		ipAddress string
		count     int
	}

	var linearr []arrstruct

	for ip, count := range logmap {
		linearr = append(linearr, arrstruct{ip, count})
	}

	sort.Slice(linearr, func(i, j int) bool {
		return linearr[i].count > linearr[j].count
	})

	for _, line := range linearr {
		fmt.Println(line.ipAddress, line.count)
	}
}
