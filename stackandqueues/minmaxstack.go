package main

import "fmt"

type MinMaxStack struct {
	stack       []int
	minMaxStack []Entry
}

type Entry struct {
	min int
	max int
}

func NewMinMaxStack() *MinMaxStack {
	return &MinMaxStack{}
}

func (s *MinMaxStack) Peek() int {
	return s.stack[len(s.stack)-1]
}

func (s *MinMaxStack) pop() int {
	s.minMaxStack = s.minMaxStack[:len(s.minMaxStack)-1]
	out := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return out
}

func (s *MinMaxStack) Push(number int) {
	newMinmax := Entry{min: number, max: number}
	if len(s.minMaxStack) > 0 {
		lastminmax := s.minMaxStack[len(s.minMaxStack)-1]
		newMinmax.min = min(lastminmax.min, number)
		newMinmax.max = max(lastminmax.max, number)
	}
	s.minMaxStack = append(s.minMaxStack, newMinmax)
	s.stack = append(s.stack, number)
}

func (s *MinMaxStack) Min() int {
	return s.minMaxStack[len(s.minMaxStack)-1].min
}
func (s *MinMaxStack) Max() int {
	return s.minMaxStack[len(s.minMaxStack)-1].max
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println("Stack and Queues")

	newMinMaxStack := NewMinMaxStack()
	newMinMaxStack.Push(5)
	fmt.Println(newMinMaxStack.Peek())
	fmt.Println(newMinMaxStack.Min())
	fmt.Println(newMinMaxStack.Max())
	newMinMaxStack.Push(2)
	fmt.Println(newMinMaxStack.Peek())
	fmt.Println(newMinMaxStack.Min())
	fmt.Println(newMinMaxStack.Max())
	newMinMaxStack.Push(12)
	fmt.Println(newMinMaxStack.Peek())
	fmt.Println(newMinMaxStack.Min())
	fmt.Println(newMinMaxStack.Max())
	newMinMaxStack.Push(7)
	fmt.Println(newMinMaxStack.Peek())
	fmt.Println(newMinMaxStack.Min())
	fmt.Println(newMinMaxStack.Max())
}
