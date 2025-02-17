package main

import "fmt"

func SortNums(nums []uint) {
	for {
		k := 0
		for i := 0; i < len(nums)-1; i++ {
			if nums[i] > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
				k++
			}
		}
		if k == 0 {
			break
		}
	}
}

func main() {
	array := []uint{490, 741, 88, 1, 10, 7, 234, 2234, 64, 12, 778, 21234, 4345, 45673, 23, 5, 78, 2, 1, 5}
	SortNums(array)
	fmt.Println(array)
}
