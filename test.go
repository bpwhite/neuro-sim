package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	var x float64
	x = 20.0
	fmt.Println(x)
	fmt.Printf("x is of type %T\n", x)

	//var a int = 0
	for a := 0; a <= 110; a++ {
		if a >= 100 {
			fmt.Printf("a is: %d\n", a)
		}
	}

}
