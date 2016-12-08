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

type alphabet struct {
	Name       int
	letter_box letter
}

type Language struct {
	Items []alphabet
}

type letter struct {
	letter_str byte
	function   int64
	sign       int64
}

func (box *Language) AddItem(item alphabet) []alphabet {
	box.Items = append(box.Items, item)
	return box.Items
}

func spawn_language(length int, langlen int64) (box Language) {
	// create 2d slice to hold language info

	// generate random string, this is pathway dna
	ran_string, err := GenerateRandomString(32)
	if err != nil {
		// Serve an appropriately vague error to the
		// user, but log the details internally.
		fmt.Println("Something terrible has happened.")
	}

	// loop through each letter of the language and assign
	// a mathematic function class
	r := ran_string
	items := []alphabet{}
	box = Language{items}

	for i := 0; i < length; i++ {
		fr := gen_cryp_num(langlen)
		sr := gen_cryp_num(2)

		// 	letter1 := letter{letter_str: 2, function: 1, sign: 1}
		l1 := letter{letter_str: r[i], function: fr, sign: sr}
		item1 := alphabet{Name: i, letter_box: l1}
		box.AddItem(item1)

		//fmt.Printf("%d %c %d %d\n", i, r[i], fr, sr)
	}
	fmt.Println("Generated language of length: ", length)
	fmt.Println("Language string: ", ran_string)

	return
}

func spawn_neuron(in_lang Language) {
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

	fmt.Println("Acquiring language...")
	lang_len := len(in_lang.Items)
	fmt.Println("Acquired language of length: ", lang_len)

	fmt.Printf("%+v\n", in_lang.Items[0].letter_box.letter_str)
	for i := 0; i < 3; i++ {
		n1.pathways = append(n1.pathways, spawn_pathway(12))
	}

	//fmt.Println("Pathway1: ", n1.pathways[0])

}

func spawn_pathway(length int) (pathway_sl [][]uint8) {

	//fmt.Println("Generated pathway with DNA string of length: ", length)

	// loop through the dna string and perform operations
	for i := 0; i < length; i++ {
		// print status of string
		//fmt.Println("[", i, "]", ran_string[i])
		// create array to hold row
		//row := []uint8{ran_string[i], 1}
		// add row to slice
		//pathway_sl = append(pathway_sl, row)

	}
	//fmt.Println(reflect.TypeOf(pathway_sl[1][0]))

	return
}
