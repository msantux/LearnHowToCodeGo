package main

import (
	"fmt"
	"slices"
	"strings"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	for i, u := range users {
		slices.Sort(u.Sayings)

		fmt.Println("User", i)
		fmt.Println("\t", u.First, u.Last, u.Age)
		for _, v := range u.Sayings {
			fmt.Println("\t\t", v)
		}
		fmt.Println()
	}

	fmt.Println("--- BY AGE ---")
	slices.SortFunc(users, func(a user, b user) int {
		return a.Age - b.Age
	})

	for i, u := range users {
		fmt.Println("User", i)
		fmt.Println("\t", u.First, u.Last, u.Age)
		for _, v := range u.Sayings {
			fmt.Println("\t\t", v)
		}
		fmt.Println()
	}

	fmt.Println("--- BY LAST ---")
	slices.SortFunc(users, func(a user, b user) int {
		return strings.Compare(a.Last, b.Last)
	})

	for i, u := range users {
		fmt.Println("User", i)
		fmt.Println("\t", u.First, u.Last, u.Age)
		for _, v := range u.Sayings {
			fmt.Println("\t\t", v)
		}
		fmt.Println()
	}

}
