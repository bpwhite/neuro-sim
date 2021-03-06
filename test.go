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

		ran := gen_cryp_num(7)
		eval := 0.0

		rsign := gen_cryp_num(2)
		sign := 1.0
		//fmt.Println(rsign)
		if rsign == 0 {
			sign = 1.0
		} else if rsign == 1 {
			sign = -1.0
		}

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
		case ran == 5:
			eval = math.Log(num)
			fmt.Printf(" log(x)=%f\t ", eval)
		case ran == 6:
			eval = math.Log10(num)
			fmt.Printf(" log10(x)=%f\t ", eval)
			//case ran == 7:
			//	eval = math.Exp(num)
			//	fmt.Printf("exp(x)=%f\t ", eval)
		}

		num = eval*sign + num
		//num = num * sign
		fmt.Printf(" num: %f\n", num)
	}
	fmt.Println("num: ", num)
}
