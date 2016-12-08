package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Beginning neural simulation...")

	language := spawn_language(26, 3)
	fmt.Println(language)

	spawn_neuron(language)

	num := 10.0

	for i := 0; i < 25; i++ {

		ran := gen_cryp_num(5)
		eval := 0.0

		switch true {
		case ran == 0:
			eval = eval + num
			fmt.Printf(" x+y=%f\t", eval)
		case ran == 1:
			eval = math.Pow(num, 2.0)
			fmt.Printf(" x^2=%f\t", eval)
		case ran == 2:
			eval = math.Pow(num, 3.0)
			fmt.Printf(" x^3=%f\t", eval)
		case ran == 3:
			eval = math.Pow(num, 4.0)
			fmt.Printf(" x^4=%f\t", eval)
		case ran == 4:
			eval = 1 / num
			fmt.Printf(" 1/x=%f\t", eval)
		}

		num = eval + num
		fmt.Printf(" num: %f\n", num)
	}
	fmt.Println("num: ", num)
}
