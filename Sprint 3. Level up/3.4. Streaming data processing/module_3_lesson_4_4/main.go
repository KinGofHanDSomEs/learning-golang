package main

import "fmt"

func MultiplyPipeline(inputNums ...[]int) int {
	var nums []int
	for _, r := range inputNums {
		for _, n := range r {
			nums = append(nums, n)
		}
	}
	return Multiply(Filter(NumbersGen(nums...)))
}

func NumbersGen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()
	return out
}

func Filter(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			if n > 0 {
				out <- n
			}
		}
	}()
	return out
}

func Multiply(in <-chan int) int {
	res := 1
	for n := range in {
		res *= n
	}
	return res
}

func main() {
	fmt.Println(MultiplyPipeline([]int{-1, 2, 3}, []int{-1, -2, -3}, []int{1, 2, -3}))
}
