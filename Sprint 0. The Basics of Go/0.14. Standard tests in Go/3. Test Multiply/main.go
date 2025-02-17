package main

import (
	"testing"
)

func TestMultiply(t *testing.T) {
	got := Multiply(2, 3)
	expected := 6

	if got != expected {
		t.Fatalf("Multiply(2, 3) = %d, want %d", got, expected)
	}
}
