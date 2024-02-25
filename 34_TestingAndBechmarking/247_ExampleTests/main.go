package main

import (
	"LearnHowToCodeGo/34_TestingAndBechmarking/247_ExampleTests/myMath"
	"fmt"
)

func main() {
	fmt.Println("2 + 3 =", myMath.MySum(2, 3))
	fmt.Println("4 + 7 =", myMath.MySum(4, 7))
	fmt.Println("5 + 9 =", myMath.MySum(5, 9))
}
