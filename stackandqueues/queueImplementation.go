package main

import "fmt"

type MyQueue struct {
	data []int
}

func NewQueue() *MyQueue {
	return &MyQueue{}
}

func (mq *MyQueue) Peek() int {
	if len(mq.data) == 0 {
		fmt.Println("Empty Queue")
		return -1
	}
	return mq.data[0]
}
func (mq *MyQueue) Enqueue(value int) {
	mq.data = append(mq.data, value)
}

func (mq *MyQueue) Dequeue() int {
	if len(mq.data) == 0 {
		fmt.Println("Empty Queue")
		return -1
	}
	value := mq.data[0]
	mq.data = mq.data[1:]
	return value
}

func main() {
	queue := NewQueue()
	queue.Enqueue(5)
	fmt.Println("Peek:", queue.Peek()) // Should print 5
	queue.Enqueue(2)
	fmt.Println("Peek:", queue.Peek()) // Should print 5
	queue.Enqueue(12)
	fmt.Println("Peek:", queue.Peek()) // Should print 5
	queue.Enqueue(7)
	fmt.Println("Peek:", queue.Peek()) // Should print 5

	fmt.Println("Dequeue:", queue.Dequeue()) // Should print 5
	fmt.Println("Peek:", queue.Peek())       // Should print 2

	fmt.Println("Dequeue:", queue.Dequeue()) // Should print 2
	fmt.Println("Peek:", queue.Peek())       // Should print 12

	fmt.Println("Dequeue:", queue.Dequeue()) // Should print 12
	fmt.Println("Peek:", queue.Peek())       // Should print 7

	fmt.Println("Dequeue:", queue.Dequeue()) // Should print 7

	// Trying to dequeue from an empty queue
	fmt.Println("Dequeue:", queue.Dequeue()) // Should print "Queue is empty" and return -1

}
