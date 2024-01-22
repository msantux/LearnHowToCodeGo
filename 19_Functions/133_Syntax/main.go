package main

import "fmt"

func main() {
	foo()
	bar("Todd")
	fmt.Println(aloha("Todd"))

	s, age := dogYears("Todd", 47)
	fmt.Println(age, "years.", s)
}

// no parameter - no returns
func foo() {
	fmt.Println("I'm from foo")
}

// one parameter - no return
func bar(s string) {
	fmt.Println("My name is", s)
}

// one parameter - one return
func aloha(s string) string {
	return fmt.Sprint("They call me Mr.", s)
}

// two parameters - two return
func dogYears(name string, age int) (string, int) {
	age *= 7

	return fmt.Sprint(name, " is this old in dog years."), age
}
