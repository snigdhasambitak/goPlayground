package main

import "fmt"

type Courses struct {
	CourseID      int
	Prerequisites []*Courses
}

func NewCourse(courseId int) *Courses {
	return &Courses{CourseID: courseId}
}

func (c *Courses) AddCourses(prerequisites ...*Courses) {
	c.Prerequisites = append(c.Prerequisites, prerequisites...)
}

func findCycleinCourses(c *Courses, stack map[int]bool, visited map[int]bool) bool {
	if stack[c.CourseID] {
		return true
	}

	if visited[c.CourseID] {
		return false
	}

	stack[c.CourseID] = true
	visited[c.CourseID] = true

	for _, child := range c.Prerequisites {
		if findCycleinCourses(child, stack, visited) {
			return true
		}
	}
	stack[c.CourseID] = false
	return false
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	// create graph nodes
	nodes := make(map[int]*Courses)
	for i := 0; i < numCourses; i++ {
		nodes[i] = NewCourse(i)
	}

	for _, prereq := range prerequisites {
		course, pre := prereq[0], prereq[1]
		nodes[course].AddCourses(nodes[pre])
	}

	stack := make(map[int]bool)
	visited := make(map[int]bool)

	for i := 0; i < numCourses; i++ {
		if !visited[i] {
			if findCycleinCourses(nodes[i], stack, visited) {
				return false
			}
		}

	}
	return true

}

func main() {
	// Example 1
	numCourses := 2
	prerequisites := [][]int{{1, 0}}
	fmt.Println(canFinish(numCourses, prerequisites)) // Output: true

	// Example 2
	numCourses = 2
	prerequisites = [][]int{{1, 0}, {0, 1}}
	fmt.Println(canFinish(numCourses, prerequisites)) // Output: false

}
