package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Println("--------- A Tour of Go ---------")
	fmt.Println()

	fmt.Println("**** Step 11 - Basic Types ****")
	fmt.Printf("Type: %T - Value %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T - Value %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T - Value %v\n", z, z)
	fmt.Println()

	fmt.Println("**** Step 12 - Zero Values ****")
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
	fmt.Println()

	fmt.Println("**** Step 13 - Type Conversions ****")
	var x, y int = 3, 4
	var f2 float64 = math.Sqrt(float64(x*x + y*y))
	var z2 uint = uint(f2)
	fmt.Printf("x= %v - y= %v - f2= %v - z2= %v", x, y, f2, z2)
	fmt.Println()
}
