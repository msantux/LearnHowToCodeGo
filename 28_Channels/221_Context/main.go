package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	ctx := context.Background()

	fmt.Println("context: \t", ctx)
	fmt.Println("context err: \t", ctx.Err())
	fmt.Printf("contet type: \t%T\n", ctx)
	fmt.Println("----------------")

	ctx, cancel := context.WithCancel(ctx)
	fmt.Println("context: \t", ctx)
	fmt.Println("context err: \t", ctx.Err())
	fmt.Printf("contet type: \t%T\n", ctx)
	fmt.Println("cancel: \t\t", cancel)
	fmt.Printf("cancel type: \t%T\n", cancel)
	fmt.Println("----------------")

	cancel()
	fmt.Println("context: \t", ctx)
	fmt.Println("context err: \t", ctx.Err())
	fmt.Printf("contet type: \t%T\n", ctx)
	fmt.Println("cancel: \t\t", cancel)
	fmt.Printf("cancel type: \t%T\n", cancel)
	fmt.Println("----------------")

	example_01()
}

func example_01() {
	ctx, cancel := context.WithCancel(context.Background())

	fmt.Println("error check 1:", ctx.Err())
	fmt.Println("num go routines:", runtime.NumGoroutine())

	go func() {
		n := 0
		for {
			select {
			case <-ctx.Done():
				return
			default:
				n++
				time.Sleep(time.Millisecond * 200)
				fmt.Println("Working", n)
			}
		}
	}()

	time.Sleep(time.Second * 3)
	fmt.Println("error check 2:", ctx.Err())
	fmt.Println("num go routines:", runtime.NumGoroutine())

	fmt.Println("About to cancel context")
	cancel()
	fmt.Println("Cancelled context")

	time.Sleep(time.Second * 2)
	fmt.Println("error check 3:", ctx.Err())
	fmt.Println("num go routines:", runtime.NumGoroutine())
}
