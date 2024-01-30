package main

import "fmt"

func main() {
	x := sum([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	fmt.Println("x =", x)

}

func sum(ii []int) (total int) {
	total = 0

	for _, v := range ii {
		total += v
	}

	return
}
