package main

import (
	"cmp"
	"fmt"
	"slices"
)

type person struct {
	first string
	age   int
}

func main() {
	p1 := person{
		first: "James",
		age:   32,
	}

	p2 := person{
		first: "Moneypenny",
		age:   27,
	}

	p3 := person{
		first: "Q",
		age:   64,
	}

	p4 := person{
		first: "M",
		age:   56,
	}

	people := []person{p1, p2, p3, p4}

	fmt.Println(people)
	slices.SortFunc(people, func(a, b person) int {
		//if n := cmp.Compare(a.first, b.first); n != 0 {
		//	return n
		//}

		return cmp.Compare(a.age, b.age)
	})
	fmt.Println("--- Sorted By Age ---")
	fmt.Println(people)

	slices.SortFunc(people, func(a, b person) int {
		if n := cmp.Compare(a.first, b.first); n != 0 {
			return n
		}

		return cmp.Compare(a.age, b.age)
	})

	fmt.Println("--- Sorted By Name and then by Age ---")
	fmt.Println(people)
}
