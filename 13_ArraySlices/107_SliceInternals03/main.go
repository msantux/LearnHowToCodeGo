package main

import (
	"fmt"
	"sort"
)

func main() {
	n := []float64{3, 1, 4, 2}

	fmt.Println(medianOne(n))
	fmt.Println("after medianOne", n)

	n2 := []float64{3, 1, 4, 2}
	fmt.Println(medianTwo(n2))
	fmt.Println("after medianTwo", n2)
}

// uses the same underlying array
// everything is passed by value in go
// the value is being passed into this func
// and then assigned to x
func medianOne(x []float64) float64 {
	sort.Float64s(x)
	i := len(x) / 2

	if len(x)%2 == 1 {
		return x[i/2]
	}

	return (x[i-1] + x[i]) / 2
}

func medianTwo(x []float64) float64 {
	n := make([]float64, len(x))
	copy(n, x)

	sort.Float64s(n)
	i := len(n) / 2

	if len(n)%2 == 1 {
		return n[i/2]
	}

	return (n[i-1] + n[i]) / 2

}
