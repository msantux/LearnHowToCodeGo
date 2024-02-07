package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := "password123"

	sb, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println(s)
	fmt.Println(sb)

	loginPword := "password1234"

	err = bcrypt.CompareHashAndPassword(sb, []byte(loginPword))
	if err != nil {
		fmt.Println("YOU CAN'T LOGIN.\nError:", err)
		return
	}

	fmt.Println("You're logged in!")
}
