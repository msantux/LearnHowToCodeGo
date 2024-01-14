package main

import (
	"fmt"
)

func main() {
	x := 40
	y := 5
	fmt.Printf("x = %v \t y = %v\n", x, y)

	// LOOP
	// for statement

	for i := 0; i < 5; i++ {
		fmt.Printf("Counting to five: %v \t first foor loop\n", i)
	}

	for y < 10 {
		fmt.Printf("y = %v \t\t\t second foor loop\n", y)
		y++
	}

	// break
	// takes you out of the loop
	for {
		fmt.Printf("y = %v \t\t\t third foor loop\n", y)
		if y > 20 {
			break
		}
		y++
	}

	// continue
	// takes to the next iteration
	for i := 0; i < 20; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println("couting even numbers:", i)
	}

}
