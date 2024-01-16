package main

import "fmt"

func main() {
	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}

	age := m["James"]
	fmt.Println("The age of James is", age)

	age = m["Q"]
	fmt.Println("The age of Q is", age)

	if age, ok := m["Q"]; ok {
		fmt.Println(age)
	} else {
		fmt.Println("'Q' is not a valid key for map 'm'")
	}
}
