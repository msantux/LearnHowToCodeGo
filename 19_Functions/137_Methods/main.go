package main

import "fmt"

type person struct {
	first string
}

func (p person) speak() {
	fmt.Println("I am", p.first)
}

func main() {
	p1 := person{
		first: "James",
	}
	p1.speak()

	p2 := person{
		first: "Jenny",
	}
	p2.speak()

}
