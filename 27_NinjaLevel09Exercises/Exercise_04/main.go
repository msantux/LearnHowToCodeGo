package main

import (
	"fmt"
	"sync"
)

func main() {
	counter := 0

	gr := 500

	var mu sync.Mutex
	var wg sync.WaitGroup
	wg.Add(gr)

	for i := 0; i < gr; i++ {
		go func() {
			mu.Lock()
			v := counter
			v++
			counter = v
			fmt.Println("counter =", counter)
			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("final counter =", counter)
}
