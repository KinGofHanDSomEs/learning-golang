package main

import (
	"testing"
)

func TestContains(t *testing.T) {
	cases := []struct {
		arr      []int
		sq       int
		expected bool
	}{
		{[]int{1, 2, 3, 4}, 2, true},
		{[]int{1, 2, 3, 4}, 5, false},
		{[]int{-1, -2, -3, -4}, -2, true},
		{[]int{-1, -2, -3, -4}, -5, false},
		{[]int{}, 1, false},
	}
	for _, cs := range cases {
		t.Run("test", func(t *testing.T) {
			got := Contains(cs.arr, cs.sq)
			if got != cs.expected {
				t.Fatalf("expected: %v, got: %v", cs.expected, got)
			}
		})
	}
}
