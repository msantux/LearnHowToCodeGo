package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
	width  float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.length * s.width
}

func (c circle) area() float64 {
	return math.Pow(c.radius, 2) * math.Pi
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println("Area = ", s.area())
}

func main() {
	c := circle{
		radius: 4,
	}
	fmt.Println("Circle")
	info(c)

	s := square{
		length: 5,
		width:  8,
	}
	fmt.Println("Square")
	info(s)
}
