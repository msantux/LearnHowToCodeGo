package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Printf("Iteration %2d: ", i)
		if x := rand.Intn(5); x == 3 {
			fmt.Printf("x is 3")
		}
		fmt.Println()
	}
}
