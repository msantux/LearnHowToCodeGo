package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("--")
		for j := 0; j < 5; j++ {
			fmt.Printf("Outter loop: %v \t Inner loop: %v\n", i, j)
		}
	}

}
