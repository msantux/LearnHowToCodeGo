package main

import (
	"fmt"
	"math/rand"
)

func main() {
	c := gen()

	receive(c)

	fmt.Println("About to exit.")
}

func gen() <-chan int {
	c := make(chan int)
	const numRoutines = 10

	for i := 0; i < numRoutines; i++ {
		go func() {
			for j := 0; j < 10; j++ {
				c <- rand.Intn(1000)
			}
		}()
	}

	return c
}

func receive(c <-chan int) {
	for i := 0; i < 100; i++ {
		fmt.Println(i, "->", <-c)
	}
}
