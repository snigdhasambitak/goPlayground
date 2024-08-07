package main

import (
	"fmt"
)

// Graph represents the graph using an adjacency list
type Graph struct {
	edges map[string][]string
}

// NewGraph initializes a new graph
func NewGraph() *Graph {
	return &Graph{edges: make(map[string][]string)}
}

// AddEdge adds an edge between two nodes
func (g *Graph) AddEdge(from, to string) {
	g.edges[from] = append(g.edges[from], to)
}

// BFS performs a breadth-first search to find the shortest path
func (g *Graph) BFS(start, goal string) []string {
	// Initialize the queue with the start node
	queue := [][]string{{start}}
	visited := map[string]bool{start: true}

	// Iterate until the queue is empty
	for len(queue) > 0 {
		// Get the first path from the queue
		path := queue[0]
		queue = queue[1:]

		// Get the last node in the path
		node := path[len(path)-1]

		// If the goal node is found, return the path
		if node == goal {
			return path
		}

		// Iterate over all neighbors of the current node
		for _, neighbor := range g.edges[node] {
			if !visited[neighbor] {
				visited[neighbor] = true
				newPath := append([]string{}, path...)
				newPath = append(newPath, neighbor)
				queue = append(queue, newPath)
			}
		}
	}

	// Return an empty slice if the goal is not reachable
	return []string{}
}

func main() {
	// Create a new graph
	graph := NewGraph()

	// Add edges (flights)
	graph.AddEdge("Paris", "London")
	graph.AddEdge("Paris", "Amsterdam")
	graph.AddEdge("London", "Berlin")
	graph.AddEdge("Amsterdam", "Berlin")
	graph.AddEdge("Berlin", "Rome")
	graph.AddEdge("Amsterdam", "Rome")

	// Find the shortest path from Paris to Rome
	path := graph.BFS("Paris", "Rome")

	// Print the path
	fmt.Println("Shortest path from Paris to Rome:", path)
}
