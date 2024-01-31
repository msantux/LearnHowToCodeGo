package main

import "fmt"

func main() {
	num := 42

	fmt.Printf("Type of numPtr: %T\n", &num)
	fmt.Printf("Value of numPtr: %v\n", &num)

	y := &num
	fmt.Printf("Type of y: %T\n", y)
	fmt.Printf("Value of y: %v\n", y)
	fmt.Println(*y)

	*y = 43
	fmt.Println(num)
}
