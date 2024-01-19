package main

import "fmt"

func main() {
	am := map[string]int{
		"Todd":   42,
		"Henry":  16,
		"Padget": 14,
	}

	fmt.Println("The age of Henry was", am["Henry"])

	fmt.Println(am)
	fmt.Printf("%#v\n", am)
	fmt.Println("Length =", len(am))

	an := make(map[string]int)
	an["Lucas"] = 28
	an["Steph"] = 37

	fmt.Println(an)
	fmt.Printf("%#v\n", an)
	fmt.Println("Length =", len(an))

}
