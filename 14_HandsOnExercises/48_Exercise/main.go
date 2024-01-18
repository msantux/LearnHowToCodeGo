package main

import (
	"fmt"
	"slices"
)

func main() {
	xs := [][]string{}

	xs = slices.Insert(xs, len(xs), []string{"James", "Bond", "Shaken, not stirred"})
	xs = slices.Insert(xs, len(xs), []string{"Miss", "Moneypenny", "I'm 008"})

	for i, l := range xs {
		fmt.Printf("Record: %v: ", i)
		for j, c := range l {
			fmt.Printf("Data: %v - %-15v ", j, c)
		}
		fmt.Printf("\n")
	}

}
