package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(250)
	fmt.Printf("x = %v.\t", x)

	if x <= 100 {
		fmt.Println("Between 0 and 100.")
	} else if x > 100 && x <= 200 {
		fmt.Println("Between 101 and 200.")
	} else {
		fmt.Println("Between 201 and 250.")
	}

	for i := 0; i < 20; i++ {
		fmt.Printf("%v - rand.Intn(3) = %v\n", i, rand.Intn(3))
	}
}
