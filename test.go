package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")

	test1 := linear(.2, 55.0, 3.0)
	fmt.Printf("Test: %f\n", test1)

	spawn_neuron()

}
