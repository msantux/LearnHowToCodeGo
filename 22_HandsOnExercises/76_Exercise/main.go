package main

import "fmt"

type dog struct {
	first string
}

func (d dog) run() {
	fmt.Println("My name is", d.first, "and I'm running.")
}

func (d dog) walk() {
	fmt.Println("My name is", d.first, "and I'm walking.")
}

type youngin interface {
	walk()
	run()
}

func youngRun(y youngin) {
	y.walk()
	y.run()
}

func main() {
	d1 := dog{
		first: "Henry",
	}
	youngRun(d1)

	d2 := dog{
		first: "Padget",
	}
	youngRun(d2)
	d2 = d2.changeName("Rover")
	youngRun(d2)
}

func (d dog) changeName(s string) dog {
	d.first = s
	return d
}
