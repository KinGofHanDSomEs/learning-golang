package main

import "testing"

func TestReverseString(t *testing.T) {
	row, expected := "hello", "olleh"
	t.Run("TestReverseString", func(t *testing.T) {
		got := ReverseString(row)
		if expected != got {
			t.Fatalf("expected: %v, got: %v", expected, got)
		}
	})
}
