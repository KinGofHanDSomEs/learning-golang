package main

import "fmt"

func MaxExpressionValue(nums []int) int { // nums[s] — nums[r] + nums[q] — nums[p]
	firstArr, secondArr, thirdArr, fourthArr := make([]int, len(nums)+1), make([]int, len(nums)), make([]int, len(nums)-1), make([]int, len(nums)-2)
	for i := len(nums) - 1; i >= 0; i-- {
		firstArr[i] = max(firstArr[i+1], nums[i])
	}
	for i := len(nums) - 2; i >= 0; i-- {
		secondArr[i] = max(secondArr[i+1], firstArr[i+1]-nums[i])
	}
	for i := len(nums) - 3; i >= 0; i-- {
		thirdArr[i] = max(thirdArr[i+1], secondArr[i+1]+nums[i])
	}
	for i := len(nums) - 4; i >= 0; i-- {
		fourthArr[i] = max(fourthArr[i+1], thirdArr[i+1]-nums[i])
	}
	return fourthArr[0]
}

func main() {
	nums := []int{3, 9, 10, 1, 30, 40}
	res := MaxExpressionValue(nums)
	fmt.Println(res) // 46 (поскольку 40 – 1 + 10 – 3 - максимально)
}
