package main

import "fmt"

type engine struct {
	eletric bool
}

type vehicle struct {
	engine
	make  string
	model string
	doors int
	color string
}

func main() {
	v1 := vehicle{
		engine: engine{eletric: false},
		make:   "Nissan",
		model:  "Kicks",
		doors:  4,
		color:  "blue",
	}

	v2 := vehicle{
		engine: engine{eletric: true},
		make:   "Byd",
		model:  "Dolphin",
		doors:  4,
		color:  "Black",
	}

	fmt.Println(v1)
	fmt.Println(v2)

	fmt.Println(v1.make, v1.model, "is eletric?", v1.eletric)
	fmt.Println(v2.make, v2.model, "is eletric?", v2.eletric)
}
