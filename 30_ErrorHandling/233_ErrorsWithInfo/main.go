package main

import (
	"errors"
	"fmt"
	"log"
	"math"
)

func main() {
	//example_01()
	//example_02()
	//example_03()
	//example_04()
	example_05()
}

type norgateMathError struct {
	lat  string
	long string
	err  error
}

func (n norgateMathError) Error() string {
	return fmt.Sprintf("a norgateMathError occured: %v %v %v", n.lat, n.long, n.err)
}

func example_01() {
	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

var ErrNorgateMath = errors.New("norgate math: square root of negative number")

func example_02() {
	fmt.Printf("%T\n", ErrNorgateMath)
	_, err := sqrt2(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func example_03() {
	_, err := sqrt3(-10.23)
	if err != nil {
		log.Fatalln(err)
	}
}

func example_04() {
	_, err := sqrt4(-10.23)
	if err != nil {
		log.Fatalln(err)
	}
}

func example_05() {
	_, err := sqrt5(-10.23)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("norgate math: square root of negative number")
	}

	return math.Sqrt(f), nil
}

func sqrt2(f float64) (float64, error) {
	if f < 0 {
		return 0, ErrNorgateMath
	}

	return math.Sqrt(f), nil
}

func sqrt3(f float64) (float64, error) {
	if f < 0 {
		return 0, fmt.Errorf("norgate math: square root of negative number: %v", f)
	}

	return math.Sqrt(f), nil
}

func sqrt4(f float64) (float64, error) {
	if f < 0 {
		ErrNorgateMath2 := fmt.Errorf("norgate math: square root of negative number: %v", f)
		return 0, ErrNorgateMath2
	}

	return math.Sqrt(f), nil
}

func sqrt5(f float64) (float64, error) {
	if f < 0 {
		nme := fmt.Errorf("norgate math redux: square root of negative number: %v", f)
		return 0, norgateMathError{"50.2289 N", "99.4656 W", nme}
	}

	return math.Sqrt(f), nil
}
