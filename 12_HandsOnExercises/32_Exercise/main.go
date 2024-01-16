package main

import "fmt"

func main() {
	i := 20

	for {
		if i < 0 {
			break
		}

		fmt.Printf("Iteration %2d\n", i)

		i--
	}
}
