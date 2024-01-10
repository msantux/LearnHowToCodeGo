package main

import (
	"fmt"
	"math"
)

func main() {
	var i int8 = math.MaxInt8
	fmt.Printf("%T max value is %v\n", i, i)
	var j uint8 = math.MaxUint8
	fmt.Printf("%T max value is %v\n", j, j)
}
