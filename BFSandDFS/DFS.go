package main

import "fmt"

type Nodes struct {
	name   string
	places []*Nodes
}

func newnode(node string) *Nodes {
	return &Nodes{name: node}
}

func (n *Nodes) addchild(places ...*Nodes) {
	n.places = append(n.places, places...)
}

func (n *Nodes) DFS(array []string) []string {
	array = append(array, n.name)
	for _, child := range n.places {
		array = child.DFS(array)
	}
	return array
}

func main() {
	nodeA := newnode("A")
	nodeB := newnode("B")
	nodeC := newnode("C")
	nodeD := newnode("D")
	nodeE := newnode("E")
	nodeF := newnode("F")
	nodeG := newnode("G")
	nodeH := newnode("H")
	nodeI := newnode("I")
	nodeJ := newnode("J")
	nodeK := newnode("K")

	nodeA.addchild(nodeB, nodeC, nodeD)
	nodeB.addchild(nodeE, nodeF)
	nodeD.addchild(nodeG, nodeH)
	nodeF.addchild(nodeI, nodeJ)
	nodeG.addchild(nodeK)

	traverse := nodeA.DFS([]string{})
	fmt.Println(traverse)

}
