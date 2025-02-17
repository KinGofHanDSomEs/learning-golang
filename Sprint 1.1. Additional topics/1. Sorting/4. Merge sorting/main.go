package main

import (
	"fmt"
	"slices"
)

func SortAndMerge(left, right []int) []int {
	result := left[:]
	for _, elem := range right {
		result = append(result, elem)
	}
	slices.Sort(result)
	return result
}

func main() {
	left := []int{4, 1, 5, 0}
	right := []int{-1, 4, 5, 10}
	fmt.Println(SortAndMerge(left, right))
}
