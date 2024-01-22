package main

import "fmt"

func main() {
	fmt.Printf("add: %T\n", add)
	fmt.Printf("subtract: %T\n", subtract)
	fmt.Printf("doMath: %T\n", doMath)

	fmt.Println(doMath(42, 16, add))
	fmt.Println(doMath(42, 16, subtract))
}

func doMath(a, b int, f func(int, int) int) int {
	return f(a, b)
}

func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}
