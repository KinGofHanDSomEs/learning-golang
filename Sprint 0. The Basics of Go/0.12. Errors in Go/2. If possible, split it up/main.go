package main

import "fmt"

func DivideIntegers(a, b int) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero is not allowed")
	}
	return float64(a / b), nil
}
