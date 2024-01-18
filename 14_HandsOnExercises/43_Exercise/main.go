package main

import "fmt"

func main() {
	xi := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	for i, v := range xi {
		fmt.Printf("Index: %2v \t Value: %v \t Type: %T\n", i, v, v)
	}

	fmt.Printf("Array xi: %#v \t Type: %T\n", xi, xi)
}
