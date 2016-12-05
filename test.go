package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println("Hello, World!")

	s := person{name: "Sam", age: 50, test: 123}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)

	for a := 0; a <= 110; a++ {
		if a >= 100 {
			fmt.Printf("a is: %d\n", a)
		}
	}

}
