package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 1; i <= 42; i++ {
		fmt.Printf("Iteration %2d: ", i)

		x := rand.Intn(5)

		switch x {
		case 0:
			fmt.Println("x values is 0")
		case 1:
			fmt.Println("x values is 1")
		case 2:
			fmt.Println("x values is 2")
		case 3:
			fmt.Println("x values is 3")
		case 4:
			fmt.Println("x values is 4")
		default:
			fmt.Println("x out of the rage [0,5)")
		}
	}
}
