package main

import "fmt"

func Filter[T any](slice []T, predicate func(T) bool) []T {
	var result []T
	for _, v := range slice {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(Filter(nums, func(n int) bool { return n%2 == 0 }))
}