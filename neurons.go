// neurons.go
package main

import (
	"fmt"
)

type neuron struct {
	name     string
	age      int
	pathways []pathway
}

type pathway struct {
	letter string
	length int
}

func linear(m float32, x float32, b float32) float32 {
	// y = mx + b
	return m*x + b
}

func spawn_neuron() {
	ran_string, err := GenerateRandomString(32)
	if err != nil {
		// Serve an appropriately vague error to the
		// user, but log the details internally.
	}
	n1 := neuron{name: ran_string, age: 0}

	fmt.Println(n1.name)
	fmt.Println(n1.age)

}
