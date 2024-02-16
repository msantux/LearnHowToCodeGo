package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	counter := 0

	gr := 500

	var wg sync.WaitGroup
	wg.Add(gr)

	for i := 0; i < gr; i++ {
		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
			fmt.Println("counter =", counter)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("final counter =", counter)
}
