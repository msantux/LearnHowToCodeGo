package main

import "fmt"

func main() {
	a := func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
	}

	a()
}
