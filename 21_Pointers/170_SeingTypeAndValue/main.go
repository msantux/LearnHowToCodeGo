package main

import "fmt"

func main() {
	num := 42

	fmt.Printf("Type of numPtr: %T\n", &num)
	fmt.Printf("Value of numPtr: %v\n", &num)

	str := "James"

	fmt.Printf("Type of strPtr: %T\n", &str)
	fmt.Printf("Value of strPtr: %v\n", &str)
}
