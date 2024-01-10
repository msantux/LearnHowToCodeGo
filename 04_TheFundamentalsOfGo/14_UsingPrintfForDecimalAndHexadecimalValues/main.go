package main

import (
	"fmt"
)

func main() {
	adams := 42
	fmt.Printf("42 as binary, %b \n", adams)
	fmt.Printf("42 as hexadecimal, %x \n", adams)

	// print these values as both binary & hexadecimal
	a, b, c, d, e, f := 0, 1, 2, 3, 4, 5

	fmt.Printf("decimal: %d\t  binary: %b\t hexadecimal: %#X\n", a, a, a)
	fmt.Printf("decimal: %d\t  binary: %b\t hexadecimal: %#X\n", b, b, b)
	fmt.Printf("decimal: %d\t  binary: %b\t hexadecimal: %#X\n", c, c, c)
	fmt.Printf("decimal: %d\t  binary: %b\t hexadecimal: %#X\n", d, d, d)
	fmt.Printf("decimal: %d\t  binary: %b\t hexadecimal: %#X\n", e, e, e)
	fmt.Printf("decimal: %d\t  binary: %b\t hexadecimal: %#X\n", f, f, f)
}
