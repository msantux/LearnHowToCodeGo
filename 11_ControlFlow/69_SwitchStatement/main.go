package main

import (
	"fmt"
	"math/rand"
)

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

	if z := 2 * rand.Intn(x); z >= x {
		fmt.Printf("Z is %v and this is GREATER THAN OR EQUAL to X which is %v\n", z, x)
	} else {
		fmt.Printf("Z is %v and this is LESS THAN to X which is %v\n", z, x)
	}

	switch {
	case x < 42:
		fmt.Println("x is LESS THAN 42")
	case x == 42:
		fmt.Println("x is EQUAL TO 42")
	case x > 42:
		fmt.Println("x is GREATER THAN 42")
	default:
		fmt.Println("This is the default case for x")
	}

	switch x {
	case 40:
		fmt.Println("x is 40")
	case 41:
		fmt.Println("x is 41")
	case 42:
		fmt.Println("x is 42")
	case 43:
		fmt.Println("x is 43")
	default:
		fmt.Println("This is the default case for x")
	}

	switch x {
	case 40:
		fmt.Println("x is 40")
		fallthrough
	case 41:
		fmt.Println("Printed because of fallthrough: x is 41")
	case 42:
		fmt.Println("Printed because of fallthrough: x is 42")
	case 43:
		fmt.Println("Printed because of fallthrough: x is 43")
	default:
		fmt.Println("Printed because of fallthrough: This is the default case for x")
	}

	switch x {
	case 40:
		fmt.Println("x is 40")
		fallthrough
	case 41:
		fmt.Println("Printed because of fallthrough: x is 41")
		fallthrough
	case 42:
		fmt.Println("Printed because of fallthrough: x is 42")
		fallthrough
	case 43:
		fmt.Println("Printed because of fallthrough: x is 43")
		fallthrough
	default:
		fmt.Println("Printed because of fallthrough: This is the default case for x")
	}

}
