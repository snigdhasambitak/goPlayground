package main

import "fmt"

type Node struct {
	name     string
	children []*Node
}

func newNode(name string) *Node {
	return &Node{name: name}
}

func (n *Node) addChildNode(children ...*Node) {
	n.children = append(n.children, children...)
}

func (n *Node) DFS(array []string) []string {
	array = append(array, n.name)

	for _, child := range n.children {
		array = child.DFS(array)
	}
	return array
}

func (n *Node) BFS(array []string) []string {

	queue := []*Node{n}
	for len(queue) > 0 {
		current := queue[0]
		array = append(array, current.name)
		queue = queue[1:]
		for _, child := range current.children {
			queue = append(queue, child)
		}
	}
	return array
}

func findcyclic(n *Node, stack map[string]bool, visited map[string]bool) (string, bool) {
	if stack[n.name] {
		return n.name, true
	}
	if visited[n.name] {
		return n.name, false
	}
	stack[n.name] = true
	visited[n.name] = true

	for _, child := range n.children {
		if cyclic_node, ok := findcyclic(child, stack, visited); ok {
			return cyclic_node, true
		}
	}
	stack[n.name] = false
	return " ", false
}

func main() {
	nodeA := newNode("A")
	nodeB := newNode("B")
	nodeC := newNode("C")
	nodeD := newNode("D")
	nodeE := newNode("E")
	nodeF := newNode("F")
	nodeG := newNode("G")
	nodeH := newNode("H")
	nodeI := newNode("I")
	nodeJ := newNode("J")
	nodeK := newNode("K")

	nodeA.addChildNode(nodeB, nodeC, nodeD)
	nodeB.addChildNode(nodeE, nodeF)
	nodeD.addChildNode(nodeG, nodeH)
	nodeF.addChildNode(nodeI, nodeJ)
	nodeG.addChildNode(nodeK)

	traverseDFS := nodeA.DFS([]string{})
	fmt.Println("DFS ", traverseDFS)

	traverseBFS := nodeA.BFS([]string{})
	fmt.Println("BFS ", traverseBFS)

	nodeH.addChildNode(nodeA) // cyclic

	stack := map[string]bool{}
	visited := map[string]bool{}

	if cyclicNode, found := findcyclic(nodeA, stack, visited); found {
		fmt.Println("Graph contains a cycle involving node:", cyclicNode)
	} else {
		fmt.Println("Graph does not contain a cycle.")
	}

}
