package main

import "fmt"

func CutCable(prices []int, length int) int {
	dp := make([]int, length+1)
	for i := range dp {
		dp[i] = 0
	}
	for i := 1; i <= length; i++ {
		for j := 1; j <= i; j++ {
			if j < len(prices) {
				dp[i] = max(dp[i], prices[j]+dp[i-j])
			}
		}
	}
	return dp[length]
}

func main() {
	type testCase struct {
		prices []int
		length int
		cost   int
	}
	testCases := []testCase{
		{
			prices: []int{0, 1, 5, 8, 9, 10, 17, 17, 20},
			length: 8,
			cost:   22,
		},
		{
			prices: []int{0, 3, 5, 6, 7, 10, 12},
			length: 6,
			cost:   18,
		},
	}
	for _, tc := range testCases {
		res := CutCable(tc.prices, tc.length)
		fmt.Println(tc.cost, res)
	}
}
