package main

import "fmt"

type person struct {
	first string
	last  string
	favIC []string
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
		favIC: []string{"Chocolate", "Banana", "Passion fruit with mango and guava"},
	}

	p2 := person{
		first: "Jenny",
		last:  "Moneypenny",
		favIC: []string{"Strawberry", "Chocolate"},
	}

	fmt.Printf("Name: %v %v - Prefered ice cream flavours: ", p1.first, p1.last)

	for i, v := range p1.favIC {
		fmt.Printf("%2d - %v ", i, v)
	}
	fmt.Println()

	fmt.Printf("Name: %v %v - Prefered ice cream flavours: ", p2.first, p2.last)

	for i, v := range p2.favIC {
		fmt.Printf("%2d - %v ", i, v)
	}
	fmt.Println()
}
