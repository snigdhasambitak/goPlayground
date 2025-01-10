package main

import "fmt"

// Function to get a list of all visited cities without duplicates
func getVisitedCities(cityPairs [][]string) []string {
	visitedCities := make(map[string]struct{})

	for _, pair := range cityPairs {
		source, destination := pair[0], pair[1]
		visitedCities[source] = struct{}{}
		visitedCities[destination] = struct{}{}
	}

	uniqueCities := make([]string, 0, len(visitedCities))
	for city := range visitedCities {
		uniqueCities = append(uniqueCities, city)
	}

	return uniqueCities
}

func main() {
	cityPairs := [][]string{
		{"New York", "Los Angeles"},
		{"Chicago", "Houston"},
		{"New York", "Miami"},
		{"Los Angeles", "Chicago"},
		{"Miami", "Houston"},
	}

	visitedCities := getVisitedCities(cityPairs)
	fmt.Println("Visited Cities:", visitedCities)
}
