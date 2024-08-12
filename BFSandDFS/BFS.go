package main

import "fmt"

type BFS struct {
	Name     string
	Children []*BFS
}

func NewNodeBFS(name string) *BFS {
	return &BFS{Name: name}
}

func (n *BFS) addChildBFS(children ...*BFS) {
	n.Children = append(n.Children, children...)
}

func (n *BFS) BreadthFirstSearch(array []string) []string {
	queue := []*BFS{n}
	for len(queue) > 0 {
		current := queue[0]
		array = append(array, current.Name)
		queue = queue[1:]
		for _, child := range current.Children {
			queue = append(queue, child)
		}
	}
	return array
}

func main() {

	nodeA := NewNodeBFS("A")
	nodeB := NewNodeBFS("B")
	nodeC := NewNodeBFS("C")
	nodeD := NewNodeBFS("D")
	nodeE := NewNodeBFS("E")
	nodeF := NewNodeBFS("F")
	nodeG := NewNodeBFS("G")
	nodeH := NewNodeBFS("H")
	nodeI := NewNodeBFS("I")
	nodeJ := NewNodeBFS("J")
	nodeK := NewNodeBFS("K")

	nodeA.addChildBFS(nodeB, nodeC, nodeD)
	nodeB.addChildBFS(nodeE, nodeF)
	nodeD.addChildBFS(nodeG, nodeH)
	nodeF.addChildBFS(nodeI, nodeJ)
	nodeG.addChildBFS(nodeK)

	traveseBFS := nodeA.BreadthFirstSearch([]string{})
	fmt.Println(traveseBFS)

}
