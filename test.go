package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")

	test1 := linear(.2, 55.0, 3.0)
	fmt.Printf("Test: %f\n", test1)

	spawn_neuron()

	for a := 0; a <= 110; a++ {
		if a >= 100 {
			//fmt.Printf("a is: %d\n", a)
		}
	}
}
