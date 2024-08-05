package main

import "fmt"

type Nodes struct {
	value int
	next  *Nodes
}

func main() {

	node1 := &Nodes{value: 1}
	node2 := &Nodes{value: 2}
	node3 := &Nodes{value: 3}

	node1.next = node2
	node2.next = node3

	currentNode := node1

	for currentNode != nil {
		fmt.Println(currentNode.value, " ")
		currentNode = currentNode.next
	}
	
}
