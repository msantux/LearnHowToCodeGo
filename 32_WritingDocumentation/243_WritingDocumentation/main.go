package main

import (
	"LearnHowToCodeGo/32_WritingDocumentation/243_WritingDocumentation/mymath"
	"fmt"
)

func main() {
	fmt.Println("2 + 3 =", mymath.Sum(2, 3))
	fmt.Println("4 + 7 =", mymath.Sum(4, 7))
	fmt.Println("5 + 9 =", mymath.Sum(5, 9))
}