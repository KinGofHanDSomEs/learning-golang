package main

func Process(nums []int) chan int {
	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		if i+1 > len(nums) {
			break
		}
		ch <- nums[i]
	}
	return ch
}
