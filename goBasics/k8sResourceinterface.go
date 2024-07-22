package main

import "fmt"

type K8sresources interface {
	Describe() string
}

type Deployment struct {
	Name     string
	Replicas int
}

type Service struct {
	Name string
	Port int
}

func (d Deployment) Describe() string {
	return fmt.Sprintf("Deployment name: %s, Replicas: %d", d.Name, d.Replicas)
}

func (s Service) Describe() string {
	return fmt.Sprintf("Service name: %s, Port: %d", s.Name, s.Port)
}

func PrintResources(resources K8sresources) {
	fmt.Println(resources.Describe())
}

func main() {
	d := Deployment{"nginx", 2}
	s := Service{"nginx-service", 80}
	PrintResources(d)
	PrintResources(s)
}
