package main

import "fmt"

func main() {
	fmt.Println(printSquare(square, 3))
}

func printSquare(f func(int) int, n int) string {
	return fmt.Sprintf("The number %v squared is %v", n, f(n))
}

func square(n int) int {
	return n * n
}
