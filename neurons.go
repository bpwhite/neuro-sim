// neurons.go
package main

type neuron struct {
	name     string
	age      int
	pathways []pathway
}

type pathway struct {
	letter string
	length int
}

func linear(x int, b int) int {
	return x + b
}
