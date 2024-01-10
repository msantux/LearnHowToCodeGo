package main

import "fmt"

const Pi = 3.14

func main() {
	fmt.Println("--------- A Tour of Go ---------")
	fmt.Println()

	fmt.Println("**** Step 14 - Type Inference ****")
	v := 42
	fmt.Printf("Type: %T - Value %v\n", v, v)
	fmt.Println()

	fmt.Println("**** Step 15 - Constants ****")
	const world = "世界"
	fmt.Println("Hello", world)
	fmt.Println("Happy", Pi, "day!")

	const Truth = true
	fmt.Println("Go rules?", Truth)
	fmt.Println()
}
