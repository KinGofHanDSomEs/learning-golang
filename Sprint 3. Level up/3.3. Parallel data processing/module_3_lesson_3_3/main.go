package main

import (
	"errors"
	"fmt"
	"sync"
)

var (
	err = errors.New("invalid data")
)

func ProcessSum(summer func(arr []int, result chan<- int), nums []int, chunkSize int) (int, error) {
	if chunkSize < 1 {
		return 0, err
	}
	resultChan, wg, result, i := make(chan int, len(nums)/chunkSize+1), sync.WaitGroup{}, 0, 0
	for i < len(nums) {
		var slice []int
		if i > len(nums)-chunkSize {
			slice = nums[i:]
		} else {
			slice = nums[i : i+chunkSize]
		}
		wg.Add(1)
		go func() {
			defer wg.Done()
			summer(slice, resultChan)
		}()
		i += chunkSize
	}
	go func() {
		wg.Wait()
		close(resultChan)
	}()
	for r := range resultChan {
		result += r
	}
	return result, nil
}

func SumChunk(arr []int, result chan<- int) {
	var res int
	for _, num := range arr {
		res += num
	}
	result <- res
}

func main() {
	result1, err1 := ProcessSum(SumChunk, []int{3, 5, 3, 6, 6}, 2)
	result2, err2 := ProcessSum(SumChunk, []int{3, 5, 3, 6, 6, 3, 5, 3, 6, 6}, 3)
	result3, err3 := ProcessSum(SumChunk, []int{3, 5, 3, 6, 6}, -10)
	fmt.Println(result1, err1)
	fmt.Println(result2, err2)
	fmt.Println(result3, err3)
	fmt.Println(ProcessSum(SumChunk, []int{3, 5, 3, 6, 6}, 1))
	fmt.Println(ProcessSum(SumChunk, []int{1, 2, 3}, 4))
}
