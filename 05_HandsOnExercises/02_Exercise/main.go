package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Hello, Marcelo!")
	fmt.Println("Welcome to VSCode!")
	fmt.Println("The time is:", time.Now())
	fmt.Println()

	fmt.Println("--------- A Tour of Go ---------")
	fmt.Println()

	fmt.Println("**** Step 01 - Packages ****")
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Println()

	fmt.Println("**** Step 02 - Imports ****")
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	fmt.Println()

	fmt.Println("**** Step 03 - Exported names ****")
	/* 	In Go, a name is exported if it begins with a capital letter.
	For example, Pi is an exported name, which is exported from the math package.
	pi does not start with a capital letter, so it is not exported.
	When importing a package, you can refer only to its exported names.
	Any "unexported" names are not accessible from outside the package.
	*/
	fmt.Println("Pi:", math.Pi)
	fmt.Println()
}
