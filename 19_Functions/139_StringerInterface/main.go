package main

import (
	"fmt"
	"log"
	"strconv"
)

type book struct {
	title string
}

func (b book) String() string {
	return fmt.Sprint("The title of the book is ", b.title)
}

type count int

func (c count) String() string {
	return fmt.Sprint("This is the number: ", strconv.Itoa(int(c)))
}

func logInfo(s fmt.Stringer) {
	log.Println("Log from 138 -", s.String())
}

func main() {
	b := book{
		title: "West with the night",
	}

	var c count = 42

	logInfo(b)
	logInfo(c)
}