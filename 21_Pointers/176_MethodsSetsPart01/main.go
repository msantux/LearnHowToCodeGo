package main

import "fmt"

type dog struct {
	first string
}

func (d dog) walk() {
	fmt.Println("My name is", d.first, "and I'm walking.")
}

func (d *dog) run() {
	d.first = "Rover"
	fmt.Println("My name is", d.first, "and I'm running.")
}

func main() {
	d1 := dog{
		first: "Henry",
	}
	d1.walk()
	d1.run()

	d2 := &dog{
		first: "Padget",
	}
	d2.walk()
	d2.run()

}
