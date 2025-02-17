package main

import "testing"

func TestSum(t *testing.T) {
	got := Sum(1, 2)
	expected := 3
	if got != expected {
		t.Fatalf("Sum(1, 2) = %d, want %d", got, expected)
	}
}
