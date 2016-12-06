package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")

	fmt.Println("Test: ", linear(5, 3))

	s := neuron{name: "Sam", age: 50}
	fmt.Println(s.name)

	sp := &s
	//fmt.Println(sp.age)

	sp.age = 51
	//fmt.Println(sp.age)

	for a := 0; a <= 110; a++ {
		if a >= 100 {
			//fmt.Printf("a is: %d\n", a)
		}
	}
}
