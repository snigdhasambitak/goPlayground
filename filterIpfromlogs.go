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
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	r, err := regexp.Compile(`\b\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}\b`)
	if err != nil {
		log.Fatal(err)
	}

	ipaddmap := make(map[string]int)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		matches := r.FindAllString(line, -1)
		for _, match := range matches {
			ipaddmap[match]++
		}

	}
	fmt.Println(ipaddmap)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// create a slice of structs to store ipaddresses and their counts

	type ipCount struct {
		Ip    string
		Count int
	}

	var ipCounts []ipCount

	for ip, count := range ipaddmap {
		ipCounts = append(ipCounts, ipCount{ip, count})
	}

	// sort the slice using an anonymous function
	sort.Slice(ipCounts, func(i, j int) bool {
		return ipCounts[i].Count > ipCounts[j].Count
	})

	for _, ipcount := range ipCounts {
		fmt.Println(ipcount.Ip, ipcount.Count)
	}

}
