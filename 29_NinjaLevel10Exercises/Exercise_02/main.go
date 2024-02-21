package main

import (
	"fmt"
)

func main() {
	step01()
	step02()
}

func step01() {
	cs := make(chan int)

	go func() {
		cs <- 42
	}()

	fmt.Println(<-cs)

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)
}

func step02() {
	c := make(chan int)

	go func(cr chan<- int) {
		cr <- 42
	}(c)
	fmt.Println(<-c)

	fmt.Printf("------\n")
	fmt.Printf("cr\t%T\n", c)
}
