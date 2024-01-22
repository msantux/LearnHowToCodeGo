package main

import "fmt"

func main() {
	foo()

	func() {
		fmt.Println("Anonymous func ran")
	}()

	func(s string) {
		fmt.Println("This is my name:", s)
	}("Todd")
}

func foo() {
	fmt.Println("Foo ran")
}
