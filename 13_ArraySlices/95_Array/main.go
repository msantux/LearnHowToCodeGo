package main

import "fmt"

func main() {
	{
		var x [5]int
		fmt.Println(x)

		x[3] = 42
		fmt.Println(x)
		fmt.Println(len(x))
	}

	{
		// declare a variable of type [7]int
		var ni [7]int

		//assign a value to index position zero
		ni[0] = 42
		fmt.Printf("%#v \t\t\t\t %T\n", ni, ni)

		// array lietral
		ni2 := [4]int{55, 56, 57, 58}
		fmt.Printf("%#v \t\t\t\t\t %T\n", ni2, ni2)

		// array literal
		// have compiler count elements
		ns := [...]string{"Chocolate", "Vanilla", "Strawberry"}
		fmt.Printf("%#v \t %T\n", ns, ns)

		// use the builtin len function
		fmt.Println("ni length =", len(ni))
		fmt.Println("ni2 length =", len(ni2))
		fmt.Println("ns length =", len(ns))
	}
}
