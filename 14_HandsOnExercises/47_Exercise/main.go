package main

import (
	"fmt"
)

func main() {
	usaStates := make([]string, 0, 50)
	fmt.Printf("Slice: length =  %v \t capacity = %v\n", len(usaStates), cap(usaStates))

	usaStates = append(
		usaStates,
		"Alabama",
		"Alaska",
		"Arizona",
		"Arkansas",
		"California",
		"Colorado",
		"Connecticut",
		"Delaware",
		"Florida",
		"Georgia",
		"Hawaii",
		"Idaho",
		"Illinois",
		"Indiana",
		"Iowa",
		"Kansas",
		"Kentucky",
		"Louisiana",
		"Maine",
		"Maryland",
		"Massachusetts",
		"Michigan",
		"Minnesota",
		"Mississippi",
		"Missouri",
		"Montana",
		"Nebraska",
		"Nevada",
		"New Hampshire",
		"New Jersey",
		"New Mexico",
		"New York",
		"North Carolina",
		"North Dakota",
		"Ohio",
		"Oklahoma",
		"Oregon",
		"Pennsylvania",
		"Rhode Island",
		"South Carolina",
		"South Dakota",
		"Tennessee",
		"Texas",
		"Utah",
		"Vermont",
		"Virginia",
		"Washington",
		"West Virginia",
		"Wisconsin",
		"Wyoming",
	)

	fmt.Printf("Slice: length =  %v \t capacity = %v\n", len(usaStates), cap(usaStates))

	for i := 0; i < len(usaStates); i++ {
		fmt.Printf("%2v - %v\n", i, usaStates[i])
	}
}
