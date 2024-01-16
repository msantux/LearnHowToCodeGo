package main

import (
	"fmt"
	"math/rand"
)

func main() {

	for j := 0; j < 100; j++ {
		fmt.Printf("%d\t", j)
	}
	fmt.Println()
	fmt.Println()

	for i := 1; i <= 100; i++ {
		x := rand.Intn(10)
		y := rand.Intn(10)

		fmt.Printf("Iteration %d: x = %v and y = %v.\t", i, x, y)

		switch {
		case x < 4 && y < 4:
			fmt.Println("x and y are both less than 4.")
		case x > 6 && y > 6:
			fmt.Println("x and y are both greater than 6.")
		case x >= 4 && x <= 6:
			fmt.Println("x is and greater than or equal to 4 and less than or equal to 6.")
		case y != 5:
			fmt.Println("y is not 5.")
		default:
			fmt.Println("None of the previous cases were met.")
		}
	}
}
