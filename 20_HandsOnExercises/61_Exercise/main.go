package main

import "fmt"

type person struct {
	first string
	age   int
}

func (p person) speak() {
	fmt.Printf("My name is %v, and I am %v years old.\n", p.first, p.age)
}

func main() {
	p := person{
		first: "James",
		age:   37,
	}

	p.speak()
}
