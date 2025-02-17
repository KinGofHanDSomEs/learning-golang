package main

func Mix(nums []int) []int {
	res := make([]int, 0, len(nums))
	for i := 0; i < len(nums)/2; i++ {
		res = append(res, nums[i], nums[i+len(nums)/2])
	}
	return res
}