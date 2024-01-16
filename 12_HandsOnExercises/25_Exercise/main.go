package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(250)
	fmt.Printf("x = %v.\t", x)

	switch {
	case x <= 100:
		fmt.Println("Between 0 and 100.")
	case x <= 200:
		fmt.Println("Between 101 and 200.")
	case x <= 250:
		fmt.Println("Between 201 and 250.")
	default:
		fmt.Println("Greater than 250.")
	}
}
