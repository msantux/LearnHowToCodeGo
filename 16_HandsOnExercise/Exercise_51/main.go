package main

import (
	"fmt"
)

func main() {
	m := map[string][]string{}

	m["bond_james"] = []string{"shaken, not stirred", "martinis", "fast cars"}
	m["moneypenny_jenny"] = []string{"intelligence", "literature", "computer_science"}
	m["dr_no"] = []string{"cats", "ice cream", "sunsets"}
	m["fleming_ian"] = []string{"stakes", "cigars", "espionage"}

	fmt.Println(mapString(m))

	fmt.Println("--------------------------")
	fmt.Println()

	delete(m, "flemming_ian")
	fmt.Println(mapString(m))

}

func mapString(m map[string][]string) string {
	s := ""
	for k, slice := range m {
		s += fmt.Sprintf("%-20v -> ", k)
		for i, v := range slice {
			if i != 0 {
				s += " // "
			}
			s += fmt.Sprintf("%v - %v", i, v)
		}
		s += "\n"
	}

	return s
}
