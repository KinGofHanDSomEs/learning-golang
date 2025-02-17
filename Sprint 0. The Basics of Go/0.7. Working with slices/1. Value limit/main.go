package main

import "errors"

func UnderLimit(nums []int, limit int, n int) ([]int, error) {
	if n <= 0 {
		return nil, errors.New("n должно быть положительным числом")
	}
	if nums == nil {
		return nil, errors.New("вы ввели пустой массив")
	}
	res := make([]int, 0, n)
	for _, num := range nums {
		if len(res) == cap(res) {
			break
		}
		if num < limit {
			res = append(res, num)
		}
	}
	return res, nil
}
