package main

import "fmt"

// Node represents an element in the linked list
type Node struct {
	value int
	next  *Node
}

// LinkedList represents the linked list
type LinkedList struct {
	head *Node
}

// Add adds a new node with the given value to the end of the list
func (ll *LinkedList) Add(value int) {
	newNode := &Node{value: value}
	if ll.head == nil {
		ll.head = newNode
		return
	}
	current := ll.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

// Remove removes the first occurrence of the given value from the list
func (ll *LinkedList) Remove(value int) {
	if ll.head == nil {
		return
	}
	if ll.head.value == value {
		ll.head = ll.head.next
		return
	}
	current := ll.head
	for current.next != nil && current.next.value != value {
		current = current.next
	}
	if current.next != nil {
		current.next = current.next.next
	}
}

// Print prints the values in the linked list
func (ll *LinkedList) Print() {
	current := ll.head
	for current != nil {
		fmt.Print(current.value, " ")
		current = current.next
	}
	fmt.Println()
}

// Size returns the number of nodes in the linked list
func (ll *LinkedList) Size() int {
	count := 0
	current := ll.head
	for current != nil {
		count++
		current = current.next
	}
	return count
}

// Search searches for a node with the given value and returns true if found
func (ll *LinkedList) Search(value int) bool {
	current := ll.head
	for current != nil {
		if current.value == value {
			return true
		}
		current = current.next
	}
	return false
}

// Reverse reverses the linked list
func (ll *LinkedList) Reverse() {
	var prev, next *Node
	current := ll.head
	for current != nil {
		next = current.next
		current.next = prev
		prev = current
		current = next
	}
	ll.head = prev
}

func main() {
	ll := LinkedList{}

	// Add nodes
	ll.Add(1)
	ll.Add(2)
	ll.Add(3)
	ll.Print() // Output: 1 2 3

	// Remove a node
	ll.Remove(2)
	ll.Print() // Output: 1 3

	// Get size
	fmt.Println("Size:", ll.Size()) // Output: Size: 2

	// Search for a value
	fmt.Println("Search 3:", ll.Search(3)) // Output: Search 3: true
	fmt.Println("Search 4:", ll.Search(4)) // Output: Search 4: false

	// Reverse the linked list
	ll.Reverse()
	ll.Print() // Output: 3 1
}
