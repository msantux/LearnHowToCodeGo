package main

import "fmt"

func main() {
	m := map[string][]string{}

	m["bond_james"] = []string{"shaken, not stirred", "martinis", "fast cars"}
	m["moneypenny_jenny"] = []string{"intelligence", "literature", "computer_science"}
	m["dr_no"] = []string{"cats", "ice cream", "sunsets"}
	m["fleming_ian"] = []string{"stakes", "cigars", "espionage"}

	for k, slice := range m {
		fmt.Printf("%-20v -> ", k)
		for i, v := range slice {
			if i != 0 {
				fmt.Printf(" // ")
			}
			fmt.Printf("%v - %v", i, v)
		}
		fmt.Printf("\n")
	}
}
