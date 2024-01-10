package main

import "fmt"

var c, python, java bool
var c2, python2, java2 = true, false, "no!"

func main() {
	fmt.Println("--------- A Tour of Go ---------")
	fmt.Println()

	fmt.Println("**** Step 08 - Variable ****")
	var i int
	fmt.Println("i: =", i, "- c =", c, "- pytho =", python, "- java =", java)
	fmt.Println()

	fmt.Println("**** Step 09 - Variables with Initializers ****")
	var j, k int = 1, 2
	fmt.Println("j =", j, "- k =", k, "- c =", c2, "- pytho =", python2, "- java =", java2)
	fmt.Println()

	fmt.Println("**** Step 10 - Short Variables Declarations ****")
	var l, m int = 1, 2
	n := 3
	c3, python3, java3 := true, false, "no!"
	fmt.Println("l =", l, "- m =", m, "- n =", n, "- c =", c3, "- pytho =", python3, "- java =", java3)
	fmt.Println()
}
