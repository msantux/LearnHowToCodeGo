package main

import "fmt"

type person struct {
	first string
}

func changeName(p person, s string) person {
	p.first = s
	return p
}

func changeNamePtr(p *person, s string) {
	p.first = s
}

func main() {
	p := person{
		first: "James",
	}
	fmt.Println(p)

	p = changeName(p, "Jenny")
	fmt.Println(p)

	changeNamePtr(&p, "James")
	fmt.Println(p)
}
