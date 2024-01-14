package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// CONCURRENCY
	// select channel

	ch1 := make(chan int)
	ch2 := make(chan int)

	// get an int64 convert to type time.Durantion, tehn use it in time.Sleep
	// Int63n returns an int64
	// type duration underlyne type is int64
	// time.Sleep takes time.Duration
	// time.Millisencond is a constant in the time package
	d1 := time.Duration(rand.Int63n(250))
	d2 := time.Duration(rand.Int63n(250))
	//fmt.Printf("%v \t %T\n", d1, d1)
	//time.Sleep(d1 * time.Millisecond)
	//fmt.Println("DONE")

	go func() {
		time.Sleep(d1 * time.Millisecond)
		ch1 <- 41
	}()

	go func() {
		time.Sleep(d2 * time.Millisecond)
		ch2 <- 42
	}()

	// A select statement chooses wich of a set of a possible send or receive operations will proceed.
	// It looks similar to a switch statement but with the cases all refereing to communication operations.
	select {
	case v1 := <-ch1:
		fmt.Println("Value from channel 1:", v1)
	case v2 := <-ch2:
		fmt.Println("Value from channel 2:", v2)
	}

}
