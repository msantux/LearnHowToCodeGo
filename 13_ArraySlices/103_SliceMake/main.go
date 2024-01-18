package main

import "fmt"

func main() {
	si := []string{"a", "b", "c"}
	fmt.Printf("si - %#v\n", si)

	xi := make([]int, 0, 10)
	fmt.Printf("xi - %#v\n", xi)
	fmt.Println("xi length =", len(xi))
	fmt.Println("xi capacity =", cap(xi))

	xi = append(xi, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Printf("x1 - %#v\n", xi)
	fmt.Println("xi length =", len(xi))
	fmt.Println("xi capacity =", cap(xi))
	fmt.Println("-------------")
	xi = append(xi, 10, 11, 12, 13)
	fmt.Printf("x1 - %#v\n", xi)
	fmt.Println("xi length =", len(xi))
	fmt.Println("xi capacity =", cap(xi))
	fmt.Println("-------------")
}
