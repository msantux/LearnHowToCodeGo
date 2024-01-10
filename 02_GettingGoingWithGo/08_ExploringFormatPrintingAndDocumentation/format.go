package main

import "fmt"

func main() {
	const name = "Kim"
	const age = 22

	// const name, age = "Kim", 22

	//fmt.Printf("%s is %d years old.\n", name, age)
	fmt.Printf("%v is %v years old.\t And the type is %T and %T.", name, age, name, age)
}
