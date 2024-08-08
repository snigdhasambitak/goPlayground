package main

import "fmt"

type MinMaxStack2 struct {
	stack  []int
	minmax []Items
}
type Items struct {
	min int
	max int
}

func NewMinMaxStack2() *MinMaxStack2 {
	return &MinMaxStack2{}
}

func (s *MinMaxStack2) peek() int {
	return s.stack[len(s.stack)-1]
}

func (s *MinMaxStack2) pop() int {
	if len(s.stack) == 0 {
		return -1
	}
	out := s.stack[len(s.stack)-1]
	s.minmax = s.minmax[:len(s.minmax)-1]
	s.stack = s.stack[:len(s.minmax)-1]
	return out
}

func (s *MinMaxStack2) push(value int) {
	newminmax := Items{value, value}
	if len(s.minmax) > 0 {
		lastminmax := s.minmax[len(s.minmax)-1]
		newminmax.min = Min(lastminmax.min, newminmax.min)
		newminmax.max = Max(lastminmax.max, newminmax.max)
	}
	s.minmax = append(s.minmax, newminmax)
	s.stack = append(s.stack, value)
}

func (s *MinMaxStack2) MinItem() int {
	return s.minmax[len(s.minmax)-1].min
}
func (s *MinMaxStack2) MaxItem() int {
	return s.minmax[len(s.minmax)-1].max
}

func Min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
func Max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

func main() {
	fmt.Println("Stack and Queues")

	newMinMaxStack := NewMinMaxStack2()
	newMinMaxStack.push(5)
	fmt.Println(newMinMaxStack.peek())
	fmt.Println(newMinMaxStack.MinItem())
	fmt.Println(newMinMaxStack.MaxItem())
	newMinMaxStack.push(2)
	fmt.Println(newMinMaxStack.peek())
	fmt.Println(newMinMaxStack.MinItem())
	fmt.Println(newMinMaxStack.MaxItem())
	newMinMaxStack.push(7)
	fmt.Println(newMinMaxStack.peek())
	fmt.Println(newMinMaxStack.MinItem())
	fmt.Println(newMinMaxStack.MaxItem())
}
