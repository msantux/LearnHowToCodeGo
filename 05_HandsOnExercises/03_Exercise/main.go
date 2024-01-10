package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println("--------- A Tour of Go ---------")
	fmt.Println()

	fmt.Println("**** Step 04 and 05 - Functions ****")
	fmt.Println("42 + 13 =", add(42, 13))
	fmt.Println()

	fmt.Println("**** Step 06 - Multiple Results ****")
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println()

	fmt.Println("**** Step 07 - Named Return Values ****")
	fmt.Println(split(17))
	fmt.Println()
}
