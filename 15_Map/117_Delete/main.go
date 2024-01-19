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
	an["George"] = 78

	fmt.Println(an)
	fmt.Printf("%#v\n", an)
	fmt.Println("Length =", len(an))

	delete(an, "George")
	fmt.Println(an)
	fmt.Printf("%#v\n", an)
	fmt.Println("Length =", len(an))

	fmt.Println("--- Acessing keys tha don't exist ---")
	delete(an, "Georgina")
	fmt.Println(an["George"])
	fmt.Println("-------------------------------------")

	if v, ok := an["Georgey"]; !ok {
		fmt.Println("Key didn't exist.")
	} else {
		fmt.Println("The value prints:", v)
	}

	for k, v := range an {
		fmt.Printf("%v age is %v\n", k, v)
	}
}
