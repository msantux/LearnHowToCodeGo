package main

import "fmt"

func main() {
	x := foo()
	y, s := bar()

	fmt.Println("foo = ", x)
	fmt.Println("bar = ", y, "and", s)
}

func foo() int {
	return 42
}

func bar() (int, string) {
	return 43, "James Bond"
}
