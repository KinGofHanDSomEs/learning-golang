package main

import (
	"fmt"
)

func Factorial(n int) (int, error) {
	if n < 0 {
		return 0, fmt.Errorf("factorial is not defined for negative numbers")
	}
	fact := 1
	for i := 1; i <= n; i++ {
		fact *= i
	}
	return fact, nil
}
