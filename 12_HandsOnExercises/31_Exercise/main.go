package main

import "fmt"

func main() {
	i := 20

	for i >= 0 {
		fmt.Printf("Iteration %2d\n", i)
		i--
	}
}
