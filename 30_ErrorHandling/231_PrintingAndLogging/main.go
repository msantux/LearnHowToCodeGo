package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	example_01()
	example_02()
	//example_03()
	//example_04()
	example_05()
}

func example_01() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		fmt.Println("err happened", err)
		log.Println("err happened", err)
		//log.Fatalln(err)
		//panic(err)
	}
}

func example_02() {
	f, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	log.SetOutput(f)

	f2, err := os.Open("no-file.txt")
	if err != nil {
		//fmt.Println("err happened", err)
		log.Println("err happened", err)
		//log.Fatalln(err)
		//panic(err)
	}
	defer f2.Close()

	fmt.Println("Check the log.txt file in the directory")

}

func example_03() {
	log.SetOutput(os.Stdout)
	_, err := os.Open("no-file.txt")
	if err != nil {
		log.Fatalln(err)
		//panic(err)
	}
}

func example_04() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		log.Panic(err)
	}
}

func example_05() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		panic(err)
	}
}
