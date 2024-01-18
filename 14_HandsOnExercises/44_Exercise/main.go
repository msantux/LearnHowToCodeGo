package main

import (
	"fmt"
)

func main() {
	xi := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	xs1 := xi[:5]
	fmt.Printf("Slice 1 - %#v\n", xs1)

	xs2 := xi[5:]
	fmt.Printf("Slice 2 - %#v\n", xs2)

	xs3 := xi[2:7]
	fmt.Printf("Slice 2 - %#v\n", xs3)

	xs4 := xi[1:6]
	fmt.Printf("Slice 2 - %#v\n", xs4)
}
