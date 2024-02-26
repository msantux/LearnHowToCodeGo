package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	y := Years(4)
	if y != 28 {
		t.Error("Expected: 28 \tGot:", y)
	}
}

func TestYearsTwo(t *testing.T) {
	y := YearsTwo(4)
	if y != 28 {
		t.Error("Expected: 28 \tGot:", y)
	}
}

func ExampleYears() {
	fmt.Println(Years(4))
	// Output:
	// 28
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(4))
	// Output:
	// 28
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(4)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(4)
	}
}
