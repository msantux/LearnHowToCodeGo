package main

import "fmt"

// value semantics
func addOne(v int) int {
	return v + 1
}

// pointer semantics
func addOneP(v *int) {
	*v += 1
}

func main() {
	// value semantics
	a := 1
	fmt.Println("a =", a)
	fmt.Println("addOne(a) =", addOne(a))
	fmt.Println("a =", a)

	// pointer semantics
	fmt.Println("--- Pointer Semantics ---")
	b := 1
	fmt.Println("b =", b)
	addOneP(&b)
	fmt.Println("b =", b)

}
