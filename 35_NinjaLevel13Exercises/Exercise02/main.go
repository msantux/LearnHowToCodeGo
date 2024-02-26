package main

import (
	"LearnHowToCodeGo/35_NinjaLevel13Exercises/Exercise02/quote"
	"LearnHowToCodeGo/35_NinjaLevel13Exercises/Exercise02/word"
	"fmt"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
