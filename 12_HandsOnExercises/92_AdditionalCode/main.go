package main

import "fmt"

func main() {
	x := "James Bond"
	if x == "James Bond" {
		fmt.Println(x)
	}

	y := "Moneypenny"
	if y == "Moneypenny" {
		fmt.Println(y)
	} else if y == "James Bond" {
		fmt.Println("BONDBONDBONDBOND", y)
	} else {
		fmt.Println("Neither.")
	}

	switch {
	case false:
		fmt.Println("Dosen't print")
	case true:
		fmt.Println("Prints")
	case true:
		fmt.Println("Dosen't print")
	}

	favSport := "surfing"
	switch favSport {
	case "skiing":
		fmt.Println("Go to the mountain!")
	case "swimming":
		fmt.Println("Go to the pool!")
	case "surfing":
		fmt.Println("Go to the Hawaii!")
	case "wingsuit flying":
		fmt.Println("Living dangerously!")
	}
}
