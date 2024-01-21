package main

import "fmt"

type person struct {
	first string
	last  string
	favIC []string
}

func main() {
	persons := map[string]person{}

	persons["bond"] = person{
		first: "James",
		last:  "Bond",
		favIC: []string{"Chocolate", "Banana", "Passion fruit with mango and guava"},
	}

	persons["moneypenny"] = person{
		first: "Jenny",
		last:  "Moneypenny",
		favIC: []string{"Strawberry", "Chocolate"},
	}

	for _, p := range persons {
		fmt.Printf("Name: %v %v - Prefered ice cream flavours: ", p.first, p.last)

		for i, v := range p.favIC {
			fmt.Printf("%2d - %v ", i+1, v)
		}
		fmt.Println()
	}
}
