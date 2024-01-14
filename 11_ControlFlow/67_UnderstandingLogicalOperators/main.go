package main

import "fmt"

func main() {
	// SEQUENCE
	fmt.Println("This is the first statement to run")
	fmt.Println("This is the second statement to run")
	x := 40 // This is the third statement to run
	y := 5  // This is the fourth statement to run
	fmt.Printf("x = %v \t y = %v\n", x, y)

	// CONDITIONAL
	// if statement
	// switch statement

	if x < 42 {
		fmt.Println("Less than the meaning of life")
	}

	if x < 42 {
		fmt.Println("Less than the meaning of life")
	} else {
		fmt.Println("Greater than, or equal to, the meaning of life")
	}

	if x < 42 {
		fmt.Println("Less than the meaning of life")
	} else if x == 42 {
		fmt.Println("Equal to, the meaning of life")
	} else {
		fmt.Println("Greater than the meaning of life")
	}

	if x < 42 && y < 42 {
		fmt.Println("Both are less than the meaning of life")
	}

	if x > 30 || x < 42 {
		fmt.Println("X is getting close to the meaning of life")
	}

	if x != 42 && y != 10 {
		fmt.Println("X is not 42 and Y is not 10")
	}

}
