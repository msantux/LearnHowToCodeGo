package main

import "fmt"

var (
	a, b, c *string
	d       *int
)

func init() {
	p := "Drop by drop, the bucket gets filled."
	q := "Persistently, patiently, you are bound to succed."
	r := "The meaning of life is..."
	n := 42
	a = &p
	b = &q
	c = &r
	d = &n
}

func main() {
	fmt.Printf("a -> \tValue: %v \tType: %T\n", a, a)
	fmt.Printf("b -> \tValue: %v \tType: %T\n", b, b)
	fmt.Printf("c -> \tValue: %v \tType: %T\n", c, c)
	fmt.Printf("d -> \tValue: %v \tType: %T\n", d, d)

	fmt.Printf("*a -> %v\n", *a)
	fmt.Printf("*b -> %v\n", *b)
	fmt.Printf("*c -> %v\n", *c)
	fmt.Printf("*d -> %v\n", *d)

}
