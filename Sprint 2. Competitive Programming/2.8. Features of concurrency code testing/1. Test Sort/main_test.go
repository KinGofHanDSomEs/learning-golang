package main

import (
	"sort"
	"testing"
)

func TestSortIntegers(t *testing.T) {
	nums := []int{1, 5, 3, 9, 4}
	res := nums[:]
	t.Run("test", func(t *testing.T) {
		t.Parallel()
		SortIntegers(nums)
		sort.Ints(res)
		if len(res) != len(nums) {
			t.Fatalf("different array lengths, case: %v, result: %v", nums, res)
		}
		for i := range nums {
			if res[i] != nums[i] {
				t.Fatalf("expected: %v, got %v", res, nums)
			}
		}
	})
}
