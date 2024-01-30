package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	x := foo(xi...)
	fmt.Println("x =", x)

	y := bar(xi)
	fmt.Println("y =", y)
}

func foo(n ...int) int {
	total := 0

	for _, v := range n {
		total += v
	}

	return total
}

func bar(n []int) int {
	total := 0

	for _, v := range n {
		total += v
	}

	return total
}
