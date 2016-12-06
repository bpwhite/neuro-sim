// neurons.go
package main

import (
	"fmt"
	//"os"
	//"reflect"
	//"crypto/rand"
	//"math/big"
)

type neuron struct {
	name     string
	age      int
	pathways [][][]uint8 // i dont know why this works but it does.
	energy   int
}

type pathway struct {
	letter string
	length int
}

func spawn_neuron() {
	ran_string, err := GenerateRandomString(12)
	if err != nil {
		// Serve an appropriately vague error to the
		// user, but log the details internally.
		fmt.Println("Something terrible has happened.")
	}
	n1 := neuron{name: ran_string, age: 0}

	fmt.Println("Neuron: ", 1)
	fmt.Println("Name: ", n1.name)
	fmt.Println("Age: ", n1.age)

	for i := 0; i < 3; i++ {
		//n1.pathways[i] = spawn_pathway(12)
		n1.pathways = append(n1.pathways, spawn_pathway(12))
		//fmt.Println(reflect.TypeOf(spawn_pathway(12)))
		//fmt.Println(spawn_pathway(12))
	}

	fmt.Println("Pathway1: ", n1.pathways[0])

}

func spawn_pathway(length int) (pathway_sl [][]uint8) {
	// create 2d slice to hold pathway info

	// generate random string, this is pathway dna
	ran_string, err := GenerateRandomString(length)
	if err != nil {
		// Serve an appropriately vague error to the
		// user, but log the details internally.
		fmt.Println("Something terrible has happened.")
	}

	fmt.Println("Generated DNA string of length: ", length)

	// loop through the dna string and perform operations
	for i := 0; i < length; i++ {
		// print status of string
		//fmt.Println("[", i, "]", ran_string[i])
		// create array to hold row
		row := []uint8{ran_string[i], 1}
		// add row to slice
		pathway_sl = append(pathway_sl, row)

	}
	//fmt.Println(reflect.TypeOf(pathway_sl[1][0]))

	return
}

func spawn_language(length int, langlen int64) (language [][]uint8) {
	// create 2d slice to hold language info

	// generate random string, this is pathway dna
	ran_string, err := GenerateRandomString(length)
	if err != nil {
		// Serve an appropriately vague error to the
		// user, but log the details internally.
		fmt.Println("Something terrible has happened.")
	}

	fmt.Println("Generated language of length: ", length)
	fmt.Println("Language string: ", ran_string)
	// loop through each letter of the language and assign
	// a mathematic function class
	for i := 0; i < length; i++ {
		fmt.Println("Selected: ", gen_cryp_num(langlen))
	}
	return
}
