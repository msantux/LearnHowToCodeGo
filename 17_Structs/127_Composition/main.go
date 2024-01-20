package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type foo int

func main() {
	var a foo = 42
	fmt.Println("a =", a)
	fmt.Printf("a type: %T\n", a)

	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	p2 := person{
		first: "Jenny",
		last:  "Moneypenny",
		age:   27,
	}

	fmt.Printf("%T\n", p1)
	fmt.Printf("%T\n", p2)
}
