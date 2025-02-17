package main

func Join(nums1, nums2 []int) []int {
	nums3 := make([]int, 0, len(nums1)+len(nums2))
	for _, num := range nums1 {
		nums3 = append(nums3, num)
	}
	for _, num := range nums2 {
		nums3 = append(nums3, num)
	}
	return nums3
}
