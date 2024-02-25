// Package myMath implements some mathematical operations
package myMath

// MySum add un unlimited number of values of type int.
func MySum(xi ...int) int {
	sum := 0

	for _, v := range xi {
		sum += v
	}

	return sum
}
