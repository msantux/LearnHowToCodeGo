package main

import (
	"fmt"
)

type customErr struct {
	info string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("ERROR customErr: %v", ce.info)
}

func main() {
	c1 := customErr{"Custom error"}
	foo(c1)
}

func foo(e error) {
	fmt.Println(e)
}
