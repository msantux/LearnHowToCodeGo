package main

import "fmt"

var d = 42

const f = 3.14

func main() {
	s := "Hello"

	fmt.Printf("d value is %v and its type is %T\n", d, d)
	fmt.Printf("f value is %v and its type is %T\n", f, f)
	fmt.Printf("s value is %#v and its type is %T\n", s, s)

}
