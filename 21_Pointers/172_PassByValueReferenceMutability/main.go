package main

import "fmt"

func main() {
	a := 42
	fmt.Println("a =", a)

	intDelta(&a)
	fmt.Println("a =", a)

	xi := []int{1, 2, 3, 4}
	fmt.Println("xi =", xi)

	sliceDelta(xi)
	fmt.Println("xi =", xi)

	m := make(map[string]int)
	m["James"] = 32
	fmt.Println("m =", m)

	mapDelta(m, "James")
	fmt.Println("m =", m)
}

func intDelta(n *int) {
	*n = 43
}

func sliceDelta(xi []int) {
	xi[0] = 99
}

func mapDelta(m map[string]int, s string) {
	m[s] = 33
}
