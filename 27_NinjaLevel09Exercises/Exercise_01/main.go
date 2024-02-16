package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(2)
	fmt.Println("Num Go Routines:", runtime.NumGoroutine())
	go func() {
		runtime.Gosched()
		fmt.Println("Hello from thing one")
		wg.Done()
	}()
	fmt.Println("Num Go Routines:", runtime.NumGoroutine())

	go func() {
		runtime.Gosched()
		fmt.Println("Hello from thing two")
		wg.Done()
	}()
	fmt.Println("Num Go Routines:", runtime.NumGoroutine())

	fmt.Println("Num Go Routines:", runtime.NumGoroutine())
	wg.Wait()

	fmt.Println("About to exit.")
}
