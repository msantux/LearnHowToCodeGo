package main

import "fmt"

type person struct {
	first string
}

func (p *person) speak() {
	fmt.Println("Hello, my name is", p.first)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		first: "James",
	}

	//saySomething(p1) dosen't work
	saySomething(&p1)
}
