package main

import "fmt"

func main() {
	// for range loop
	// data structure - slice
	xi := []int{41, 42, 43, 44, 45, 46, 47}

	for i, v := range xi {
		fmt.Println("Ranging over a slice", i, v)
	}

	// for range loop
	// data structure - map
	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}

	for k, v := range m {
		fmt.Println("Ranging over a map", k, v)
	}
}
