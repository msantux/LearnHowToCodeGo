package main

import "fmt"

func main() {
	fmt.Println("Entering first loop")
	for i := 0; i < 5; i++ {
		fmt.Printf("Starting first loop iteration %d\n", i)
		fmt.Println("Entering second loop")
		for j := 0; j < 5; j++ {
			fmt.Printf("Starting second loop iteration %d\t", j)
			fmt.Printf("i = %v \t j = %v\n", i, j)
		}
		fmt.Println("Finishing second loop iteration.")
		fmt.Println()
	}
	fmt.Println("Finishing first loop iteration.")
}
