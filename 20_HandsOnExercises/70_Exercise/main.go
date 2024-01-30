package main

import "fmt"

func main() {
	a := getFunc()

	fmt.Println(a())
}

func getFunc() func() int {
	return func() int {
		return 42
	}
}
