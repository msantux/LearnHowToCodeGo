package main

import "fmt"

func main() {
	p1 := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "James",
		friends: map[string]int{
			"Jenny Moneypenny": 27,
			"Q":                87,
			"Ian":              47,
		},
		favDrinks: []string{
			"Martini",
			"Water",
		},
	}

	fmt.Println(p1)
}
