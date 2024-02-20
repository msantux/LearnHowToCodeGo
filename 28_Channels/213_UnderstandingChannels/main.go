package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)

	b := make(chan int, 1)

	b <- 43

	fmt.Println(<-b)

	a := make(chan int, 1)

	go func() {
		a <- 44
		a <- 45
	}()

	fmt.Println(<-a)
	fmt.Println(<-a)
}
