package main

import "math"

func FindMaxKey(m map[int]int) int {
	max_key := math.MinInt
	for key, _ := range m {
		if key > max_key {
			max_key = key
		}
	}
	return max_key
}
