package main

import "fmt"

func main() {
	foo()

	a := func() {
		fmt.Println("Anonymous func ran")
	}

	b := func(s string) {
		fmt.Println("This is my name:", s)
	}

	a()
	b("Todd")
}

func foo() {
	fmt.Println("Foo ran")
}
