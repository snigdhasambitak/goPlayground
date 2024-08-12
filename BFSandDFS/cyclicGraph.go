package main

import "fmt"

type CycleGraph struct {
	Node     string
	Children []*CycleGraph
}

func NewNode(value string) *CycleGraph {
	return &CycleGraph{Node: value}
}

func (g *CycleGraph) AddEdges(childs ...*CycleGraph) {
	g.Children = append(g.Children, childs...)
}

func verifyCycle(g *CycleGraph, stack map[string]bool, visited map[string]bool) bool {

	if stack[g.Node] {
		return true
	}
	if visited[g.Node] {
		return false
	}

	stack[g.Node] = true
	visited[g.Node] = true

	for _, child := range g.Children {
		if verifyCycle(child, stack, visited) {
			return true
		}
	}
	stack[g.Node] = false
	return false
}

func main() {
	// Creating a graph with a cycle
	nodeA := NewNode("A")
	nodeB := NewNode("B")
	nodeC := NewNode("C")

	nodeA.AddEdges(nodeB)
	nodeB.AddEdges(nodeC)
	nodeC.AddEdges(nodeA) // Creates a cycle

	visited := make(map[string]bool)
	stack := make(map[string]bool)

	hasCycle := verifyCycle(nodeA, visited, stack)
	fmt.Println("Cycle detected:", hasCycle)
}
