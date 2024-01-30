package main

import "testing"

func TestAdd(t *testing.T) {
	total := add(5, 5)

	if total != 10 {
		t.Errorf("Expected %d - Found %d", 10, total)
	}
}

func TestSubtractc(t *testing.T) {
	total := subtract(5, 5)

	if total != 0 {
		t.Errorf("Expected %d - Found %d", 0, total)
	}
}
