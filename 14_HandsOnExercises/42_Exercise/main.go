package main

import (
	"fmt"
	"math/rand"
)

func main() {
	ai := [5]int{}

	for i := range ai {
		ai[i] = rand.Intn(100)
	}

	for i, v := range ai {
		fmt.Printf("Index: %v \t Value: %v \t Type: %T\n", i, v, v)
	}
}
