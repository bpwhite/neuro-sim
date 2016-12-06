// neurons.go
package main

import (
	"fmt"
	"os"
	"reflect"
)

type neuron struct {
	name     string
	age      int
	pathways []pathway
	energy   int
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
	ran_string, err := GenerateRandomString(12)
	if err != nil {
		// Serve an appropriately vague error to the
		// user, but log the details internally.
		fmt.Println("Something terrible has happened.")
	}
	n1 := neuron{name: ran_string, age: 0}

	for i := 0; i < 3; i++ {
		//n1.pathways[i] := spawn_pathway()
		spawn_pathway(12)
	}

	fmt.Println("Name: ", n1.name)
	fmt.Println("Age: ", n1.age)
	//fmt.Println("Pathway1: ", n1.pathways[0])

}

func spawn_pathway(length int) {
	ran_string, err := GenerateRandomString(length)
	if err != nil {
		// Serve an appropriately vague error to the
		// user, but log the details internally.
		fmt.Println("Something terrible has happened.")
	}
	fmt.Println(reflect.TypeOf(ran_string[1]), " : ", ran_string[1])
	fmt.Println("Length: ", length)

	pathway_sl := [][]uint8{}

	for i := 0; i < length; i++ {
		row := []uint8{ran_string[i], 1}
		pathway_sl = append(pathway_sl, row)
		//pathway_arr[i][0] = ran_string[i]
		//pathway_arr[i][1] = 1
		fmt.Println("[", i, "]", ran_string[i])
	}
	fmt.Println(pathway_sl[1][0])
	// do something with c
	//fmt.Println(c)
	//pathway_arr[_][1] := c, 1
	//}
	os.Exit(3)

	//retun ran_string
}
