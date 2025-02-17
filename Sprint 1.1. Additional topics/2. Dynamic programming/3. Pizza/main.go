package main

import (
	"math"
)

func MinPizzaCost(s, m, l, cs, cm, cl, x int) int {
	maxSize := max(s, m, l)
	dp := make([]int, x+maxSize+1)
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	for i := 1; i <= x+maxSize; i++ {
		if i >= s {
			dp[i] = min(dp[i], dp[i-s]+cs)
		}
		if i >= m {
			dp[i] = min(dp[i], dp[i-m]+cm)
		}
		if i >= l {
			dp[i] = min(dp[i], dp[i-l]+cl)
		}
	}
	minCost := math.MaxInt32
	for i := x; i <= x+maxSize; i++ {
		if dp[i] < minCost {
			minCost = dp[i]
		}
	}
	if minCost == math.MaxInt32 {
		return -1
	}
	return minCost
}
