package main

import "fmt"

type Graph struct {
	Edge map[string][]string
}

func NewGraph() *Graph {
	return &Graph{make(map[string][]string)}
}

func (g *Graph) AddEdge(from, to string) {
	g.Edge[from] = append(g.Edge[from], to)
}

func (g *Graph) FindShortestPath(start, end string) []string {

	queue := [][]string{{start}}
	visited := map[string]bool{start: true}

	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]
		node := path[len(path)-1]

		if node == end {
			return path
		}

		for _, child := range g.Edge[node] {
			if !visited[child] {
				visited[child] = true
				tempPath := append([]string{}, path...)
				tempPath = append(tempPath, child)
				queue = append(queue, tempPath)
			}
		}
	}
	return []string{}
}

func main() {
	g := NewGraph()
	// Add edges (flights)
	g.AddEdge("Paris", "London")
	g.AddEdge("Paris", "Amsterdam")
	g.AddEdge("London", "Berlin")
	g.AddEdge("Amsterdam", "Berlin")
	g.AddEdge("Berlin", "Rome")
	g.AddEdge("Amsterdam", "Rome")

	fmt.Println(g)
	fmt.Println(g.FindShortestPath("Paris", "Rome"))

}
