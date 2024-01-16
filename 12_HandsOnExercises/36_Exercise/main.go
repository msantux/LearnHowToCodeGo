package main

import "fmt"

func main() {
	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}

	for k, v := range m {
		fmt.Printf("Key: %-15s Value: %d\n", k, v)
	}
}
