package main

import "testing"

func TestParadise(t *testing.T) {
	got := Paradise("Bali")
	want := "My idea of paradise is Bali"

	if got != want {
		t.Errorf("Error - Want: %s - Get: %s", want, got)
	}
}
