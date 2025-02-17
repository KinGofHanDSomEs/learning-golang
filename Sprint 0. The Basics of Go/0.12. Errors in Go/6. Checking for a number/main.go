package main

import (
	"fmt"
	"strconv"
)

func SumTwoIntegers(a, b string) (int, error) {
	a1, ok1 := strconv.Atoi(a)
	b1, ok2 := strconv.Atoi(b)
	if ok1 != nil || ok2 != nil {
		return 0, fmt.Errorf("invalid input, please provide two integers")
	}
	return a1 + b1, nil
}
