package main

import "fmt"

func main() {
	x := foo()
	fmt.Println(x)

	y := bar()
	fmt.Println(y())

	fmt.Printf("foo: %T\n", foo)
	fmt.Printf("bar: %T\n", bar)
	fmt.Printf("y: %T\n", y)
}

func foo() int {
	return 42
}

func bar() func() int {
	return func() int {
		return 43
	}
}
