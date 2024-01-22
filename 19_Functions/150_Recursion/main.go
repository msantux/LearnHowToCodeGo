package main

import "fmt"

func main() {
	// 4! = 4 x 3 x 2 x 1
	fmt.Println(4 * 3 * 2 * 1)
	fmt.Println(factorial(4))
	fmt.Println(factLoop(4))
}

func factorial(i int) int {
	if i == 0 {
		return 1
	}

	return i * factorial(i-1)
}

func factLoop(i int) int {
	total := 1
	for i > 0 {
		total *= i
		i--
	}

	return total
}
