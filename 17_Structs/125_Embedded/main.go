package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	ltk bool
}

func main() {
	p1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   32,
		},
		ltk: true,
	}

	p2 := secretAgent{
		person: person{
			first: "Jenny",
			last:  "Moneypenny",
			age:   27,
		},
		ltk: false,
	}

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Printf("%T \t %#v\n", p1, p1)
	fmt.Printf("%T \t %#v\n", p2, p2)

	fmt.Println("Name:", p1.first, p1.last, "- Age:", p1.age, "- ltk:", p1.ltk)
	fmt.Println("p1.person =", p1.person)
}
