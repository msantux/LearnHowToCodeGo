package main

import "fmt"

func main() {
	jb := []string{"James", "Bond", "Martini", "Chocolate"}
	mp := []string{"Jenny", "Moneypenny", "Guiness", "Wolverine Tracks"}
	fmt.Println(jb)
	fmt.Println(mp)

	xxp := [][]string{jb, mp}
	fmt.Printf("%#v\n", xxp)
}
