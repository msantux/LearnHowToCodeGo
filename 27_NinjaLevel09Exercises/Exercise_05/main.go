package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	var counter atomic.Int64

	gr := 500

	var wg sync.WaitGroup
	wg.Add(gr)

	for i := 0; i < gr; i++ {
		go func() {
			runtime.Gosched()
			counter.Add(1)
			fmt.Println("counter =", counter.Load())
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("final counter =", counter.Load())
}
