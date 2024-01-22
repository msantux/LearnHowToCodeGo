package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	xb, err := readFile("poem.txt")

	if err != nil {
		log.Fatalf("readFile in main: %s", err)
	}

	fmt.Println(xb)
	fmt.Println(string(xb))
}

func readFile(fileName string) ([]byte, error) {
	xb, err := os.ReadFile(fileName)

	if err != nil {
		return nil, fmt.Errorf("Error in readFile function: %s", err)
	}

	return xb, nil
}
