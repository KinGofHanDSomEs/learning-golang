package main

func SliceCopy(nums []int) []int {
	res := make([]int, 0, len(nums))
	for _, num := range nums {
		res = append(res, num)
	}
	return res
}
